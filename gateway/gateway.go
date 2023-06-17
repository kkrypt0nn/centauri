// Package gateway handles the Gateway API of Discord
package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/centauri/gateway/packets"
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
	"io"
	"net"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"runtime"
	"strings"
	"time"
)

var ErrClientAlreadyConnected = errors.New("client is already connected to the gateway")

// Client is a client made to handle the Gateway API and REST API
type Client struct {
	Connection   *websocket.Conn
	LastSequence int64
	SessionID    string
	ResumeURL    string
	Logger       *logger.Logger
	Debug        bool

	token             string
	intents           discord.Intents
	eventHandlers     map[EventType][]EventHandler
	heartbeatInterval time.Duration
	lastHeartbeat     time.Time
	lastHeartbeatAck  time.Time
	connectionClose   chan bool
	gatewayURL        string

	restClient *rest.Client
}

var eventNameToStruct map[EventType]EventInterface

// New creates a new Gateway API and REST API client
func New(token string, intents discord.Intents) *Client {
	newLogger := logger.NewLogger()
	gatewayClient := &Client{
		Logger: newLogger,

		intents:         intents,
		gatewayURL:      endpoints.GatewayURL,
		connectionClose: make(chan bool),
	}
	gatewayClient.SetToken(token)
	restClient := &rest.Client{
		HttpClient: &rest.HttpClient{
			Client: &http.Client{
				Timeout: 20 * time.Second,
			},
			Interceptor: nil,
		},
		Logger:      newLogger,
		RateLimiter: rest.NewRateLimiter(),
		TaskManager: tasks.NewTaskManager(),
	}
	restClient.SetAuthorizationHeader(token)
	gatewayClient.SetRestClient(restClient)
	registerEventInterfaces()
	return gatewayClient
}

// SetToken sets the token header of the gateway client
func (c *Client) SetToken(token string) {
	if strings.HasPrefix(token, "Bot ") {
		c.token = strings.TrimPrefix(token, "Bot ")
	} else {
		c.token = token
	}
}

// SetRestClient sets the REST client of the gateway client
func (c *Client) SetRestClient(restClient *rest.Client) {
	c.restClient = restClient
}

// Rest returns the REST client of the gateway client
func (c *Client) Rest() *rest.Client {
	return c.restClient
}

// Login connects the client to the gateway
func (c *Client) Login() error {
	if c.Connection != nil {
		return ErrClientAlreadyConnected
	}

	c.lastHeartbeat = time.Now()

	ctx := context.Background()
	connection, response, err := websocket.Dial(ctx, c.gatewayURL, nil)
	if err != nil {
		responseBody := "<empty>"
		if response != nil && response.Body != nil {
			body, readerErr := io.ReadAll(response.Body)
			if readerErr != nil {
				return readerErr
			}
			responseBody = string(body)
		}
		return fmt.Errorf("failed connecting to the gateway (response body: %s)", responseBody)
	}

	c.Connection = connection

	// Handle hello event and wait for ready or resumed event from the gateway before handling other events
	_, message, err := connection.Reader(ctx)
	if err != nil {
		return fmt.Errorf("failed reading first event coming from the gateway: %s", err)
	}
	event, err := NewEvent(message)
	if err != nil {
		return fmt.Errorf("failed parsing an event coming from the gateway: %s", err)
	}

	// The first event coming from Discord must be an OpCode 10 (Hello) Event
	if event.OpCode != packets.OpCodeHello {
		return fmt.Errorf("expecting OpCode 10, got OpCode %d instead", event.OpCode)
	}
	if c.Debug {
		c.Logger.Info("Got HELLO event from Discord")
	}
	c.lastHeartbeatAck = time.Now()
	var eventHello EventHello
	if err = json.Unmarshal(event.Data, &eventHello); err != nil {
		return fmt.Errorf("failed unmarshalling event HELLO: %s", err)
	}
	c.heartbeatInterval = time.Duration(eventHello.HeartbeatInterval)

	if c.SessionID == "" && c.LastSequence == 0 {
		// We identify
		if err := c.identify(ctx); err != nil {
			return fmt.Errorf("failed sending IDENTIFY event: %s", err)
		}
	} else {
		// TODO: We resume
	}

	// Now we should get either a READY or RESUMED event
	_, message, err = connection.Reader(ctx)
	if err != nil {
		return fmt.Errorf("failed reading second event coming from the gateway: %s", err)
	}
	event, err = NewEvent(message)
	if err != nil {
		return fmt.Errorf("failed parsing an event coming from the gateway: %s", err)
	}
	if event.Type != string(EventTypeReady) {
		c.Logger.Warn(fmt.Sprintf("Expected a READY event but got %s instead", event.Type))
	}

	if c.Debug {
		c.Logger.Info("Got READY event from Discord")
	}

	if event.Type == string(EventTypeReady) {
		var eventReady Ready
		if err = json.Unmarshal(event.Data, &eventReady); err != nil {
			return fmt.Errorf("failed unmarshalling event READY: %s", err)
		}
		if err := c.onEvent(event, ctx); err != nil {
			return err
		}
	}

	// Start sending heartbeats and listening to events
	go c.doHeartbeat(ctx, connection)
	go c.listenForEvents(ctx, connection)

	return nil
}

func (c *Client) identify(ctx context.Context) error {
	identifyEvent := NewIdentifyEvent(c.token, ConnectionProperties{
		OS:      runtime.GOOS,
		Browser: "Centauri",
		Device:  "Centauri",
	}, c.intents)
	identifyErr := wsjson.Write(ctx, c.Connection, identifyEvent)
	if identifyErr != nil {
		return fmt.Errorf("failed sending the IDENTIFY event: %s", identifyErr)
	}
	return nil
}

func (c *Client) doHeartbeat(ctx context.Context, connection *websocket.Conn) {
	if connection == nil {
		return
	}

	ticker := time.NewTicker(c.heartbeatInterval * time.Millisecond)
	for {
		c.lastHeartbeat = time.Now()
		err := wsjson.Write(ctx, connection, NewHeartbeat(c.LastSequence))
		if err != nil || time.Now().Sub(c.lastHeartbeatAck) > (c.heartbeatInterval*5*time.Millisecond) {
			c.Logger.Error(fmt.Sprintf("Error sending heartbeat to the gateway: %s", err))
			c.Close()
			c.reconnect()
			return
		}

		if c.Debug {
			c.Logger.Info("Sent a heartbeat event")
		}

		select {
		case <-c.connectionClose:
			return
		case <-ticker.C:
		}
	}
}

// Close closes the connection to Discord
func (c *Client) Close() {
	c.CloseWithStatusCode(websocket.StatusInternalError, "the sky is falling")
}

// CloseWithStatusCode closes the connection to Discord with a custom status code
func (c *Client) CloseWithStatusCode(statusCode websocket.StatusCode, reason string) {
	_ = c.Connection.Close(statusCode, reason)
	c.Connection = nil
	c.connectionClose <- true
}

// listenForEvents will listen for gateway events and send it over to the event handler
func (c *Client) listenForEvents(ctx context.Context, connection *websocket.Conn) {
	if c.Debug {
		c.Logger.Debug("Now listening for events coming from the Gateway API")
	}
	for {
		_, message, err := connection.Reader(ctx)
		if err != nil {
			sameConnection := c.Connection == connection
			if !sameConnection {
				return
			}
			closeStatus := websocket.CloseStatus(err)
			reconnect := packets.CloseCode(closeStatus).ShouldReconnect()
			if errors.Is(err, net.ErrClosed) {
				reconnect = false
			}

			c.CloseWithStatusCode(websocket.StatusServiceRestart, "restarting...")
			if reconnect {
				c.Logger.Info("Need a gateway reconnection...")
				c.reconnect()
				return
			}

			c.Logger.Error(fmt.Sprintf("failed reading event coming from the gateway: %s", err))
			return
		}

		select {
		case <-c.connectionClose:
			return
		default:
			event, defaultErr := NewEvent(message)
			if defaultErr != nil {
				c.Logger.Error(fmt.Sprintf("failed parsing an event coming from the gateway: %s", err))
			} else {
				eventErr := c.onEvent(event, ctx)
				if eventErr != nil {
					c.Logger.Error(eventErr.Error())
				}
			}
		}
	}
}

// reconnect will reconnect to the gateway
func (c *Client) reconnect() {
	timer := time.Duration(1)
	for {
		c.Logger.Info("Trying to reconnect to the gateway")
		err := c.Login()
		if err != nil {
			if err == ErrClientAlreadyConnected {
				c.Logger.Info("Client is already connected to the gateway, restart not necessary")
				return
			}

			c.Logger.Error(fmt.Sprintf("Failed reconnecting to the gateway: %s", err))
			<-time.After(timer * time.Second)
			timer *= 2
			if timer > 600 {
				timer = 600
			}
		}

		c.Logger.Info("Successfully reconnected to the gateway")
		return
	}
}

// onEvent handles the various event opcodes
func (c *Client) onEvent(event *Event, ctx context.Context) error {
	if c.Debug {
		c.Logger.Debug(fmt.Sprintf("New Event :: OpCode: %d, Sequence %d, Type: %s, Data: %s", event.OpCode, event.Sequence, event.Type, string(event.Data)))
	}

	c.LastSequence = event.Sequence

	if event.OpCode == packets.OpCodeHeartbeat {
		if c.Debug {
			c.Logger.Info("Sending heartbeat event in response to a heartbeat request")
		}
		heartbeatEvent, err := json.Marshal(NewHeartbeat(c.LastSequence))
		if err != nil {
			return err
		}
		err = c.Connection.Write(ctx, websocket.MessageText, heartbeatEvent)
		if err != nil {
			return err
		}
		return nil
	}

	if event.OpCode == packets.OpCodeReconnect {
		if c.Debug {
			c.Logger.Info("Closing and reconnecting due to the gateway asking so")
		}
		c.CloseWithStatusCode(websocket.StatusServiceRestart, "restarting")
		c.reconnect()
		return nil
	}

	if event.OpCode == packets.OpCodeInvalidSession {
		if c.Debug {
			c.Logger.Info("Identifying once again due to the gateway asking so")
		}
		if err := c.identify(ctx); err != nil {
			return fmt.Errorf("failed sending IDENTIFY event: %s", err)
		}
		return nil
	}

	if event.OpCode == packets.OpCodeHeartbeatACK {
		if c.Debug {
			c.Logger.Info("Got a heartbeat ACK event")
		}
		c.lastHeartbeatAck = time.Now()
		return nil
	}

	if event.OpCode == packets.OpCodeDispatch {
		if h, ok := eventNameToStruct[EventType(event.Type)]; ok {
			event.Struct = h.New()
			if err := json.Unmarshal(event.Data, event.Struct); err != nil {
				return fmt.Errorf("failed unmarshalling event data of %s event: %s", event.Type, err)
			}
			c.handleEvent(event.Type, event.Struct)
		} else {
			c.Logger.Warn(fmt.Sprintf("Unknown Event :: OpCode: %d, Sequence %d, Type: %s, Data: %s", event.OpCode, event.Sequence, event.Type, string(event.Data)))
		}
	}
	return nil
}

// handleEvent handles the event and make sure it dispatches correctly
func (c *Client) handleEvent(eventName string, eventData interface{}) {
	// First some internal listeners only
	switch d := eventData.(type) {
	case *Ready:
		c.gatewayURL = d.ResumeGatewayURL
		c.SessionID = d.SessionID
	}

	// Then we dispatch them to the various user listeners
	for _, eh := range c.eventHandlers[EventType(eventName)] {
		eh.Handle(c, eventData)
	}
}

// registerEventInterfaces registers the different event names and their respective handler
func registerEventInterfaces() {
	eventNameToStruct = map[EventType]EventInterface{
		EventTypeReady:         readyEventHandler(nil),
		EventTypeMessageCreate: messageCreateEventHandler(nil),
	}
}

func (c *Client) On(eventType EventType, handlerFunction interface{}) {
	eventHandler := handlerForFunction(handlerFunction)
	if handlerFunction == nil {
		c.Logger.Error("Invalid handler function type")
		return
	}

	if c.eventHandlers == nil {
		c.eventHandlers = map[EventType][]EventHandler{}
	}
	c.eventHandlers[eventType] = append(c.eventHandlers[eventType], eventHandler)
}
