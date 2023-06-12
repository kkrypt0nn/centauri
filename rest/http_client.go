package rest

import (
	"net/http"
)

// HttpClient is a http.Client extended with a custom interceptor
type HttpClient struct {
	*http.Client
	Interceptor Interceptor
}

// SetInterceptor sets the interceptor for the HTTP client
func (c *HttpClient) SetInterceptor(interceptor Interceptor) {
	c.Interceptor = interceptor
}

// Interceptor is an interface all custom interceptor should inherit
type Interceptor interface {
	ModifyRequest(r *http.Request)
	ModifyResponse(r *http.Response)
}

// DefaultInterceptor is a default interceptor so that you do not have to add both methods in your custom interceptor
type DefaultInterceptor struct {
	Interceptor
}

// ModifyRequest is the default request modifying method from the default interceptor, which does nothing
func (i *DefaultInterceptor) ModifyRequest(_ *http.Request) {}

// ModifyResponse is the default response modifying method from the default interceptor, which does nothing
func (i *DefaultInterceptor) ModifyResponse(_ *http.Response) {}

// RequestOption is a function that will get executed before the request and can edit it. Unlike the interceptor (Interceptor), it will get executed per-request and not on all requests made by the client
type RequestOption func(r *http.Request)

// WithHeader returns a request option (RequestOption) that will set a header
func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set(key, value)
	}
}

// WithReason returns a request option (RequestOption) that will set the 'X-Audit-Log-Reason' header
func WithReason(reason string) RequestOption {
	return WithHeader("X-Audit-Log-Reason", reason)
}
