package validations

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/faceair/jio"
)

// Generic function to parse and validate request body
func parseAndValidate[T any](req *http.Request, schema *jio.ObjectSchema) (*T, error) {
	defer req.Body.Close()

	// Read request body
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("failed to read request body")
	}

	// Validate JSON payload
	if _, err = jio.ValidateJSON(&reqBody, schema); err != nil { 
		return nil, errors.New(err.Error())
	}

	// Unmarshal into the struct
	var payload T
	if err = json.Unmarshal(reqBody, &payload); err != nil {
		return nil, errors.New("failed to parse request body")
	}

	return &payload, nil
}