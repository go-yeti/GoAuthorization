// Author: James Mallon <jamesmallondev@gmail.com>
// helpers package -
package helpers

import (
	. "GoAuthorization/libs/layout"
)

// Struct type formHelper -
type formHelper struct {
	FormValidator
}

var FormHelper *formHelper

// init function - data and process initialization
func init() {
	fh := formHelper{}
	fh.Rules = map[string]string{
		"email":    `\w{2,64}@\w{2,64}\.\w{2,64}(\.\w+)?`,
		"username": `\w+`,
		"password": `[A-Za-z\d@$!%*#?&]{8,}`,
	}

	fh.Messages = map[string]string{
		"email":    "Invalid email format",
		"username": "The username must contains only alphanumeric values",
		"password": `The password must contain at least 8 characters,
                    1 uppercase character [A-Z],
                    1 lowercase character [a-z],
                    1 digit [0-9],
                    1 special character (!, $, #, etc)`,
	}
	FormHelper = &fh
}
