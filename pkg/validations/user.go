package validations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/faceair/jio"
)

type userValidation struct{}

func (u *userValidation) Login(req *http.Request) (*types.LoginUserPayload, error) {

	// Ensure the body is closed after reading it
	defer req.Body.Close()

	// Read the request body
	reqBodyJson, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Define the schema for the payload validation
	schema := jio.Object().Keys(jio.K{
		"email":    jio.String().Required(),
		"password": jio.String().Required(),
	})

	// Validate the JSON payload against the schema
	_, err = jio.ValidateJSON(&reqBodyJson, schema)
	if err != nil {
		return nil, err
	}

	// Unmarshal the validated JSON into the LoginUserPayload struct
	var loginUserPayload types.LoginUserPayload
	err = json.Unmarshal(reqBodyJson, &loginUserPayload)
	if err != nil {
		return nil, err
	}

	// Return the validated and parsed payload
	return &loginUserPayload, nil
}

func (u *userValidation) Register(req *http.Request) (*types.RegisterUserPayload, error) {

	// Ensure the body is closed after reading it
	defer req.Body.Close()

	// Read the request body
	reqBodyJson, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Define the schema for the payload validation
	schema := jio.Object().Keys(jio.K{
		"firstName": jio.String().Required(),
		"lastName":  jio.String().Required(),
		"email":    jio.String().Required(),
		"password": jio.String().Required(),
	})

	// Validate the JSON payload against the schema
	_, err = jio.ValidateJSON(&reqBodyJson, schema)
	if err != nil {
		return nil, err
	}

	// Unmarshal the validated JSON into the RegisterUserPayload struct
	var registerUserPayload types.RegisterUserPayload
	err = json.Unmarshal(reqBodyJson, &registerUserPayload)
	if err != nil {
		return nil, err
	}

	// Return the validated and parsed payload
	return &registerUserPayload, nil
}
var User = &userValidation{}
