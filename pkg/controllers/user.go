package controllers

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/pkg/lib"
	"github.com/ZUBERKHAN034/go-ecom/pkg/models"
	"github.com/ZUBERKHAN034/go-ecom/pkg/utils"
	"github.com/ZUBERKHAN034/go-ecom/pkg/validations"
)

type userController struct{}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Login godoc
//
// @Summary Login
// @Description Login
// @Tags User
// @Accept json
// @Produce json
// @Param payload body LoginUserPayload true "User Payload"
// @Success 200 {string} string "User logged in successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "User not found"
// @Failure 400 {string} string "Invalid password"
// @Router /user/login [post]
func (u *userController) Login(res http.ResponseWriter, req *http.Request) {
	var payload LoginUserPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// validate the payload
	if err := validations.User.Login(payload); err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user exists
	user := models.User.GetByEmail(payload.Email)
	if user.ID == 0 {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Email or password is incorrect")
		return
	}

	// Check if password is correct
	if !utils.ComparePassword(user.Password, payload.Password) {
		lib.SendErrorResponse(res, http.StatusBadRequest, "Invalid password")
		return
	}

	// Generate JWT token
	tokenPayload := map[string]interface{}{
		"id":    user.ID,
		"name":  user.FirstName + " " + user.LastName,
		"email": user.Email,
	}

	token, err := utils.GenerateJWT(tokenPayload)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	lib.SendSuccessResponse(res, http.StatusOK, map[string]string{"token": token})
}

// Register godoc
//
// @Summary Register
// @Description Register
// @Tags User
// @Accept json
// @Produce json
// @Param payload body RegisterUserPayload true "User Payload"
// @Success 201 {string} string "User Registered successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 400 {string} string "User already exists"
// @Failure 500 {string} string "Internal server error"
// @Router /user/register [post]
func (u *userController) Register(res http.ResponseWriter, req *http.Request) {
	var payload RegisterUserPayload

	// Parse the request body
	err := lib.ParseJSON(req, &payload)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := validations.User.Register(payload); err != nil {
		lib.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user exists
	if user := models.User.GetByEmail(payload.Email); user.ID != 0 {
		lib.SendErrorResponse(res, http.StatusBadRequest, "User already exists")
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		lib.SendErrorResponse(res, http.StatusInternalServerError, err)
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

	lib.SendErrorResponse(res, http.StatusCreated, "user Registered successfully")
}

var User = &userController{}
