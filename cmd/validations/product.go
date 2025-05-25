package validations

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/types"
	"github.com/faceair/jio"
)

type productValidation struct{}

// Product Create validation
func (p *productValidation) Create(req *http.Request) (*types.ProductPayload, error) {
	schema := jio.Object().Keys(jio.K{
		"name":        jio.String().Required(),
		"description": jio.String().Required(),
		"image":       jio.String().Required(),
		"price":       jio.Number().Required(),
		"quantity":    jio.Number().Required(),
	})

	return parseAndValidate[types.ProductPayload](req, schema)
}

var Product = &productValidation{}
