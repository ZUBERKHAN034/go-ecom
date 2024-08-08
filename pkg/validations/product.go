package validations

import (
	"encoding/json"
	"errors"

	"github.com/faceair/jio"
)

type productValidation struct{}

func (p *productValidation) CreateProduct(payload interface{}) error {
	// Define the schema for the payloads
	schema := jio.Object().Keys(jio.K{
		"name":        jio.String().Required(),
		"description": jio.String().Required(),
		"image":       jio.String().Required(),
		"price":       jio.Number().Required(),
		"quantity":    jio.Number().Required(),
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

var Product = &productValidation{}
