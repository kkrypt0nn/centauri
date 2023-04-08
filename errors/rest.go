package errors

import "errors"

var (
	Unauthorized = errors.New("the HTTP request was unauthorized, please double check the given token")
)
