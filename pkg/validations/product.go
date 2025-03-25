package validations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/faceair/jio"
)

type productValidation struct{}

func (p *productValidation) Create(req *http.Request) (*types.ProductPayload, error) {
	// Ensure the body is closed after reading it	
	defer req.Body.Close()

	// Read the request body
	reqBodyJson, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Define the schema for the payload validation
	schema := jio.Object().Keys(jio.K{
		"name":        jio.String().Required(),
		"description": jio.String().Required(),
		"image":       jio.String().Required(),
		"price":       jio.Number().Required(),
		"quantity":    jio.Number().Required(),
	})

	// Validate the JSON payload against the schema
	_, err = jio.ValidateJSON(&reqBodyJson, schema)
	if err != nil {
		return nil, err
	}

	// Unmarshal the validated JSON into the ProductPayload struct
	var productPayload types.ProductPayload
	err = json.Unmarshal(reqBodyJson, &productPayload)
	if err != nil {
		return nil, err
	}

	return &productPayload, nil
}

var Product = &productValidation{}
