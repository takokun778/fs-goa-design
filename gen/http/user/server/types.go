// Code generated by goa v3.7.5, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen fs-goa-design/design

package server

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

// NewHelloResponseBody builds the HTTP response body from the result of the
// "hello" endpoint of the "user" service.
func NewHelloResponseBody(res *user.HelloResponseBody) *HelloResponseBody {
	body := &HelloResponseBody{
		Greet: res.Greet,
	}
	return body
}

// NewHelloRequestBody builds a user service hello endpoint payload.
func NewHelloRequestBody(body *HelloRequestBody) *user.HelloRequestBody {
	v := &user.HelloRequestBody{
		Name: body.Name,
	}

	return v
}