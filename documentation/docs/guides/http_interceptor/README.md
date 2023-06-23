---
title: HTTP Interceptor
description: Centauri lets you create an HTTP Interceptor so that you can perform various actions on a response or a request.
---

:::caution

This is not intended for beginners, so use it only if you know what you're doing.

:::

Centauri lets you create an HTTP Interceptor so that you can perform various actions on a response or a request. Centauri will call the `ModifyRequest()` method of the interceptor **just before** doing the request and the `ModifyResponse()` method **just before** returning the response and handling it.

You can see an example of using an HTTP Interceptor [here](https://github.com/kkrypt0nn/centauri/blob/main/_examples/http_interceptor/main.go).

## Using interceptor

Each interceptor must have the `ModifyRequest()` and `ModifyResponse()` methods implemented, which take as argument `*http.Request` and `*http.Response` respectively.

You can set the interceptor with the `SetInterceptor()` method on the `HttpClient` field of the REST Client.

## Default interceptor

Centauri has a default interceptor (`rest.DefaultInterceptor`) which you can use that implements the `ModifyRequest()` and `ModifyResponse()` methods already so you can just override one of them.

## Example

Here is an example of an interceptor that uses the default interceptor as base so that it is only needed to specify which method we want to override. The interceptor will modify the request and change the authorization header right before the client will send the request.


```go
package main

import (
	"fmt"
	"net/http"

	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/rest"
)

// This is our custom interceptor that is based on the default interceptor
type MyInterceptor struct {
	rest.DefaultInterceptor
}

func NewCustomInterceptor() *MyInterceptor {
	return &MyInterceptor{}
}

// We just modify the request's authorization header
func (i *MyInterceptor) ModifyRequest(r *http.Request) {
	r.Header.Set("Authorization", "Bot OTHER_BOT_TOKEN")
}

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true

	// We set the interceptor
	botClient.HttpClient.SetInterceptor(NewCustomInterceptor())

	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(fmt.Sprintf("Got bot %s with ID %d", bot.Username, bot.ID))
	}
}
```