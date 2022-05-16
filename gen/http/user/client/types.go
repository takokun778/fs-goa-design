// Code generated by goa v3.7.5, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen fs-goa-design/design

package client

import (
	user "fs-goa-design/gen/user"
)

// HelloRequestBody is the type of the "user" service "hello" endpoint HTTP
// request body.
type HelloRequestBody struct {
	// user name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// HelloResponseBody is the type of the "user" service "hello" endpoint HTTP
// response body.
type HelloResponseBody struct {
	// greet result
	Greet *string `form:"greet,omitempty" json:"greet,omitempty" xml:"greet,omitempty"`
}

// NewHelloRequestBody builds the HTTP request body from the payload of the
// "hello" endpoint of the "user" service.
func NewHelloRequestBody(p *user.HelloRequestBody) *HelloRequestBody {
	body := &HelloRequestBody{
		Name: p.Name,
	}
	return body
}

// NewHelloResponseBodyOK builds a "user" service "hello" endpoint result from
// a HTTP "OK" response.
func NewHelloResponseBodyOK(body *HelloResponseBody) *user.HelloResponseBody {
	v := &user.HelloResponseBody{
		Greet: body.Greet,
	}

	return v
}