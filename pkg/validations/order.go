package validations

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/faceair/jio"
)

type orderValidation struct{}

// Order Checkout validation
func (o *orderValidation) Checkout(req *http.Request) (*types.OrderPayload, error) {
	schema := jio.Object().Keys(jio.K{
		"orderItems": jio.Array().Items(jio.Object().Keys(jio.K{
			"productId": jio.Number().Required(),
			"quantity":  jio.Number().Required(),
		})).Min(1).Required(),
	})

	return parseAndValidate[types.OrderPayload](req, schema)
}

var Order = &orderValidation{}
