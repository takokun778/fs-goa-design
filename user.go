package hellouserapi

import (
	"context"
	"fmt"
	user "fs-goa-design/gen/user"
	"log"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Hello implements hello.
func (s *usersrvc) Hello(ctx context.Context, p *user.HelloRequestBody) (res *user.HelloResponseBody, err error) {
	greet := fmt.Sprintf("hello %s", *p.Name)

	res = &user.HelloResponseBody{
		Greet: &greet,
	}
	s.logger.Print("user.hello")
	return
}
