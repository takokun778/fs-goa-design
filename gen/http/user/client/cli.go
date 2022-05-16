// Code generated by goa v3.7.5, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen fs-goa-design/design

package client

import (
	"encoding/json"
	"fmt"
	user "fs-goa-design/gen/user"
)

// BuildHelloPayload builds the payload for the user hello endpoint from CLI
// flags.
func BuildHelloPayload(userHelloBody string) (*user.HelloRequestBody, error) {
	var err error
	var body HelloRequestBody
	{
		err = json.Unmarshal([]byte(userHelloBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Et ea ea nesciunt quia eius.\"\n   }'")
		}
	}
	v := &user.HelloRequestBody{
		Name: body.Name,
	}

	return v, nil
}
