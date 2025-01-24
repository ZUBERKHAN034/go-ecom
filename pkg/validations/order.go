package validations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/faceair/jio"
)

type orderValidation struct{}

func (o *orderValidation) Checkout(req *http.Request) (*types.OrderPayload, error) {
	// Ensure the body is closed after reading it
	defer req.Body.Close()

	// Read the request body
	reqBodyJson, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	// Define the schema for the payload validation
	schema := jio.Object().Keys(jio.K{
		"orderItems": jio.Array().Items(jio.Object().Keys(jio.K{
			"productId": jio.Number().Required(),
			"quantity":  jio.Number().Required(),
		})).Min(1).Required(),
	})

	// Validate the JSON payload against the schema
	_, err = jio.ValidateJSON(&reqBodyJson, schema)
	if err != nil {
		return nil, err
	}

	// Unmarshal the validated JSON into the OrderPayload struct
	var orderPayload types.OrderPayload
	err = json.Unmarshal(reqBodyJson, &orderPayload)
	if err != nil {
		return nil, err
	}

	// Return the validated and parsed payload
	return &orderPayload, nil
}

var Order = &orderValidation{}
