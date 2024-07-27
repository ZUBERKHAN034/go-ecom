package services

import (
	"errors"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/lib"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/ZUBERKHAN034/go-ecom/pkg/utils"
)

type userService struct{}

// Method - Login
// Description - Login user
// Params - res http.ResponseWriter, req *http.Request
// Returns - any
func (u *userService) Login(res http.ResponseWriter, req *http.Request) {
	var payload types.LoginUserPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Check if user exists
	user := models.User.GetByEmail(payload.Email)
	if user.ID == 0 {
		lib.WriteError(res, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	// Check if password is correct
	if !utils.ComparePassword(user.Password, payload.Password) {
		lib.WriteError(res, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	lib.WriteSuccess(res, http.StatusOK, "user logged in successfully")
}

// Method - Register
// Description - Register user
// Params - res http.ResponseWriter, req *http.Request
// Returns - any
func (u *userService) Register(res http.ResponseWriter, req *http.Request) {
	var payload types.RegisterUserPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Check if user exists
	if user := models.User.GetByEmail(payload.Email); user.ID != 0 {
		lib.WriteError(res, http.StatusBadRequest, errors.New("user already exists"))
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		lib.WriteError(res, http.StatusInternalServerError, err)
		return
	}
	payload.Password = hashedPassword

	// Create the user
	models.User.Create((&models.UserSchema{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	}))

	lib.WriteSuccess(res, http.StatusCreated, "user Registered successfully")
}

var User = &userService{}
