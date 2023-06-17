// Package gateway handles the Gateway API of Discord
package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kkrypt0nn/centauri/endpoints"
	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
	"io"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"strings"
	"time"
)

// Client is a client made to handle the Gateway API and REST API
type Client struct {
	Connection         *websocket.Conn
	LastSequenceNumber int
	SessionID          string
	ResumeURL          string
	Logger             *logger.Logger
	Debug              bool

	token             string
	eventHandlers     map[EventName]EventHandler
	heartbeatInterval time.Duration
	lastHeartbeat     time.Time

	restClient *rest.Client
}

// New creates a new Gateway API and REST API client
func New(token string) *Client {
	newLogger := logger.NewLogger()
	gatewayClient := &Client{
		Logger: newLogger,
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
	gatewayClient.registerEventHandlers()
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

// Login connects the client to the Gateway
func (c *Client) Login() error {
	if c.Connection != nil {
		return fmt.Errorf("client is already connected to the gateway")
	}

	c.lastHeartbeat = time.Now().UTC()

	ctx := context.Background()
	connection, response, err := websocket.Dial(ctx, endpoints.GatewayURL, nil)
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

	go c.listenForEvents(ctx, connection)

	return nil
}

// Close closes the connection to Discord
func (c *Client) Close() {
	_ = c.Connection.Close(websocket.StatusInternalError, "the sky is falling")
}

// listenForEvents will listen for Gateway events and send it over to the event handler
func (c *Client) listenForEvents(ctx context.Context, connection *websocket.Conn) {
	if c.Debug {
		c.Logger.Info("Now listening for events coming from the Gateway API")
	}
	for {
		v := map[string]interface{}{}
		err := wsjson.Read(ctx, connection, &v)
		if err != nil {
			// TODO: Check if client should reconnect
			c.Logger.Error(fmt.Sprintf("error reading event coming from the Gateway: %s", err))
			return
		}

		// TODO: Check if no longer listening, then return
		select {
		default:
			c.onEvent(v)
		}
	}
}

// onEvent handles the various event opcodes
func (c *Client) onEvent(test map[string]interface{}) {
	// TODO: Handle
	jsonString, _ := json.Marshal(test)
	var e Event
	_ = json.Unmarshal(jsonString, &e)
	fmt.Println(e.OpCode)
}

// registerEventHandlers registers the different event handlers
func (c *Client) registerEventHandlers() {
	c.eventHandlers = map[EventName]EventHandler{
		EventNameReady: ReadyEventHandler(nil),
	}
}
