package validations

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/faceair/jio"
)

type userValidation struct{}

// User Login validation
func (u *userValidation) Login(req *http.Request) (*types.LoginUserPayload, error) {
	schema := jio.Object().Keys(jio.K{
		"email":    jio.String().Regex("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$").Required(),
		"password": jio.String().Regex("^[A-Za-z][A-Za-z0-9]{7,}$").Required(),
	})

	return parseAndValidate[types.LoginUserPayload](req, schema)
}

// User Register validation
func (u *userValidation) Register(req *http.Request) (*types.RegisterUserPayload, error) {
	schema := jio.Object().Keys(jio.K{
		"firstName": jio.String().Required(),
		"lastName":  jio.String().Required(),
		"email":     jio.String().Regex("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$").Required(),
		"password":  jio.String().Regex("^[A-Za-z][A-Za-z0-9]{7,}$").Required(),
	})

	return parseAndValidate[types.RegisterUserPayload](req, schema)
}

var User = &userValidation{}
