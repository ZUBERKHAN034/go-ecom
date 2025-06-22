package controllers

import (
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/app"
	"github.com/ZUBERKHAN034/go-ecom/cmd/models"
	"github.com/ZUBERKHAN034/go-ecom/cmd/types"
	"github.com/ZUBERKHAN034/go-ecom/cmd/utils"
	"github.com/ZUBERKHAN034/go-ecom/cmd/validations"
)

type userController struct{}

// Login godoc
//
// @Summary Login
// @Description Login
// @Tags User
// @Accept json
// @Produce json
// @Param payload body types.LoginUserPayload true "User Payload"
// @Success 200 {object} types.LoginResponsePayload "Login Response Payload"
// @Failure 400 {string} string "invalid request payload"
// @Failure 400 {string} string "user not exists"
// @Failure 400 {string} string "invalid password"
// @Failure 500 {string} string "failed to generate token"
// @Failure 500 {string} string "internal server error"
// @Router /users/login [post]
func (u *userController) Login(res http.ResponseWriter, req *http.Request) {

	// Validate the request payload
	user, err := validations.User.Login(req)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user exists, if not, return an error
	checkUser := models.User.GetByEmail(user.Email)
	if checkUser.ID == 0 {
		app.SendErrorResponse(res, http.StatusBadRequest, "user not exists")
		return
	}

	// Check if password is correct, if not, return an error
	if !utils.ComparePassword(checkUser.Password, user.Password) {
		app.SendErrorResponse(res, http.StatusBadRequest, "invalid password")
		return
	}

	// Generate JWT token
	tokenPayload := types.TokenPayload{
		ID:      checkUser.ID,
		Name:    checkUser.FirstName + " " + checkUser.LastName,
		Email:   checkUser.Email,
		Address: checkUser.Address,
	}

	// Generate JWT token using the utils package, if it fails, return an error
	token, err := utils.GenerateJWT(tokenPayload)
	if err != nil {
		app.SendErrorResponse(res, http.StatusInternalServerError, "failed to generate token")
		return
	}

	loginResponse := types.LoginResponsePayload{
		Token: token,
	}

	app.SendSuccessResponse(res, http.StatusOK, loginResponse)
}

// Register godoc
//
// @Summary Register
// @Description Register
// @Tags User
// @Accept json
// @Produce json
// @Param payload body types.RegisterUserPayload true "User Payload"
// @Success 201 {object} models.UserSchema "User"
// @Failure 400 {string} string "invalid request payload"
// @Failure 400 {string} string "user already exists"
// @Failure 500 {string} string "internal server error"
// @Router /users/register [post]
func (u *userController) Register(res http.ResponseWriter, req *http.Request) {

	// Validate the request payload
	user, err := validations.User.Register(req)
	if err != nil {
		app.SendErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user exists, if so, return an error
	if checkUser := models.User.GetByEmail(user.Email); checkUser.ID != 0 {
		app.SendErrorResponse(res, http.StatusBadRequest, "user already exists")
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		app.SendErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the hashed password
	user.Password = hashedPassword

	// Create the user in the database
	createdUser := models.User.Create((&models.UserSchema{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		Email:     user.Email,
		Password:  user.Password,
	}))

	app.SendSuccessResponse(res, http.StatusCreated, createdUser)
}

var User = &userController{}
