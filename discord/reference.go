package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error represents an error JSON response coming from Discord
// https://discord.com/developers/docs/reference#error-messages
type Error struct {
	Request  *http.Request  `json:"-"`
	Response *http.Response `json:"-"`

	Code    int             `json:"code"`
	Errors  json.RawMessage `json:"errors"`
	Message string          `json:"message"`
}

// NewError creates a new error coming from Discord
func NewError(body []byte, request *http.Request, response *http.Response) Error {
	var err Error
	_ = json.Unmarshal(body, &err)
	err.Request = request
	err.Response = response
	return err
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
