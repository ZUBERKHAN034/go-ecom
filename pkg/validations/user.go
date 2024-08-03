package validations

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/faceair/jio"
)

type userValidation struct{}

// Custom function to match regex
func RegexMatch(value string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func (u *userValidation) Login(payload interface{}) error {
	// Define the schema for the payloads
	schema := jio.Object().Keys(jio.K{
		"email": jio.String().Transform(func(ctx *jio.Context) {
			// Define your regex for email validation
			regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			if !RegexMatch(ctx.Value.(string), regex) {
				ctx.Abort(errors.New("field `email` value invalid"))
			}
		}),
		"password": jio.String().Min(8).Max(10).Required(),
	})

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.New("failed to marshal payload to JSON: " + err.Error())
	}

	// Validate the JSON payload against the schema
	if _, err := jio.ValidateJSON(&payloadBytes, schema); err != nil {
		return err
	}

	return nil
}

func (u *userValidation) Register(payload interface{}) error {
	// Define the schema for the payloads
	schema := jio.Object().Keys(jio.K{
		"firstName": jio.String().Required(),
		"lastName":  jio.String().Required(),
		"email": jio.String().Transform(func(ctx *jio.Context) {
			// Define your regex for email validation
			regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			if !RegexMatch(ctx.Value.(string), regex) {
				ctx.Abort(errors.New("field `email` value invalid"))
			}
		}),
		"password": jio.String().Min(8).Max(10).Required(),
	})

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errors.New("failed to marshal payload to JSON: " + err.Error())
	}

	// Validate the JSON payload against the schema
	if _, err := jio.ValidateJSON(&payloadBytes, schema); err != nil {
		return err
	}

	return nil
}

var User = &userValidation{}
