package middlewares

import (
	"context"
	"net/http"

	"github.com/ZUBERKHAN034/go-ecom/cmd/app"
	"github.com/ZUBERKHAN034/go-ecom/cmd/utils"
)

type contextAuthUser string

const authUserKey contextAuthUser = "authUser"

type UserPayload struct {
	ID      uint
	Email   string
	Name    string
	Address string
}

func AuthWithJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			app.SendErrorResponse(res, http.StatusUnauthorized, "token is missing in Authorization header")
			return
		}

		tokenPayload, err := utils.ValidateJWTToken(authHeader)
		if err != nil {
			app.SendErrorResponse(res, http.StatusUnauthorized, err.Error())
			return
		}

		user := UserPayload{
			ID:      tokenPayload.ID,
			Email:   tokenPayload.Email,
			Name:    tokenPayload.Name,
			Address: tokenPayload.Address,
		}

		ctx := context.WithValue(req.Context(), authUserKey, user)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

// GetUserPayload extracts the UserPayload from the request context
func GetAuthUser(req *http.Request) (UserPayload, bool) {
	userPayload, ok := req.Context().Value(authUserKey).(UserPayload)
	return userPayload, ok
}
