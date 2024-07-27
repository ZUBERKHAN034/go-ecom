package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/types"
	"github.com/ZUBERKHAN034/go-ecom/pkg/utils"
)

type userService struct{}

func (u *userService) Login(res http.ResponseWriter, req *http.Request) {
	var payload types.LoginUserPayload

	err := utils.ParseJSON(req, &payload)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Check if user exists
	user := models.User.GetByEmail(payload.Email)
	if user.ID == 0 {
		utils.WriteError(res, http.StatusBadRequest, errors.New("user not found"))
		return
	}

	fmt.Println("User", user)

	// Check if password is correct
	if !models.User.ComparePassword(user.Password, payload.Password) {
		utils.WriteError(res, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	utils.WriteJSON(res, http.StatusOK, user)
}

func (u *userService) Register(res http.ResponseWriter, req *http.Request) {
	var payload types.RegisterUserPayload

	err := utils.ParseJSON(req, &payload)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	// Check if user exists
	if user := models.User.GetByEmail(payload.Email); user.ID != 0 {
		utils.WriteError(res, http.StatusBadRequest, errors.New("user already exists"))
		return
	}

	// Create the user
	newUser := models.User.Create((&models.UserSchema{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	}))

	utils.WriteJSON(res, http.StatusOK, newUser)
}

var User = &userService{}
