// Code generated by goa v3.7.5, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen fs-goa-design/design

package user

import (
	"context"
)

// User Serviceドメインモデル単位くらいのニュアンス
type Service interface {
	// Hello implements hello.
	Hello(context.Context, *HelloRequestBody) (res *HelloResponseBody, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"hello"}

// HelloRequestBody is the payload type of the user service hello method.
type HelloRequestBody struct {
	// user name
	Name *string
}

// HelloResponseBody is the result type of the user service hello method.
type HelloResponseBody struct {
	// greet result
	Greet *string
}