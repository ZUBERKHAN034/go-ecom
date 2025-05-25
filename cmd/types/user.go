package types

type LoginUserPayload struct {
	// Email is the email address of the user
	// @json email
	// @example "john.doe@example.com"
	Email string `json:"email" example:"john.doe@example.com"`

	// Password is the password for the user account
	// @json password
	// @example "securepassword123"
	Password string `json:"password" example:"securepassword123"`
}

type RegisterUserPayload struct {
	// FirstName is the first name of the user
	// @json firstName
	// @example "John"
	FirstName string `json:"firstName" example:"John"`

	// LastName is the last name of the user
	// @json lastName
	// @example "Doe"
	LastName string `json:"lastName" example:"Doe"`

	// Address is the address of the user
	// @json address
	// @example "123, Street Name, City, Country"
	Address string `json:"address" example:"123, Street Name, City, Country"`

	// Email is the email address of the user
	// @json email
	// @example "john.doe@example.com"
	Email string `json:"email" example:"john.doe@example.com"`

	// Password is the password for the user account
	// @json password
	// @example "securepassword123"
	Password string `json:"password" example:"securepassword123"`
}

type TokenPayload struct {
	// ID is the unique identifier for the user
	// @json id
	// @example 1
	ID uint `json:"id" example:"1"`

	// Name is the full name of the user
	// @json name
	// @example "John Doe"

	Name string `json:"name" example:"John Doe"`
	// Email is the email address of the user
	// @json email
	// @example "john.doe@example.com"
	Email string `json:"email" example:"john.doe@example.com"`

	// Address is the address of the user
	// @json address
	// @example "123, Street Name, City, Country"
	Address string `json:"address" example:"123, Street Name, City, Country"`
}

type LoginResponsePayload struct {

	// Token is the JWT token string returned after successful login
	// @json token
	// @example "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJqb2huLmRvZUBleGFtcGxlLmNvbSIsImFkZHJlc3MiOiIxMjMsIFN0cmVldCBOYW1lLCBDaXR5LCBDb3VudHJ5In0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkpvaG4gRG9lIiwiZW1haWwiOiJqb2huLmRvZUBleGFtcGxlLmNvbSIsImFkZHJlc3MiOiIxMjMsIFN0cmVldCBOYW1lLCBDaXR5LCBDb3VudHJ5In0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}
