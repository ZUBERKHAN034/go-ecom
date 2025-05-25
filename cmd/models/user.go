package models

import (
	"time"

	"gorm.io/gorm"
)

// UserSchema represents a user in the system
// @Description User details
type UserSchema struct {
	// ID is the unique identifier for the user
	// @json id
	// @example 1
	ID uint `json:"id" example:"1"`

	// CreatedAt represents the timestamp when the user was created
	// @json createdAt
	// @example "2024-03-19T12:00:00Z"
	CreatedAt time.Time `json:"createdAt" example:"2024-03-19T12:00:00Z"`

	// UpdatedAt represents the last updated timestamp of the user
	// @json updatedAt
	// @example "2024-03-19T12:30:00Z"
	UpdatedAt time.Time `json:"updatedAt" example:"2024-03-19T12:30:00Z"`

	// DeletedAt represents the soft delete timestamp
	// @json deletedAt
	// @example null
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggerignore:"true"`

	// FirstName of the user
	// @json firstName
	// @example "John"
	FirstName string `json:"firstName" example:"John"`

	// LastName of the user
	// @json lastName
	// @example "Doe"
	LastName string `json:"lastName" example:"Doe"`

	// Address of the user
	// @json address
	// @example "123, Street Name, City, Country"
	Address string `json:"address" example:"123, Street Name, City, Country"`

	// Email of the user
	// @json email
	// @example "john.doe@example.com"
	Email string `json:"email" example:"john.doe@example.com"`

	// Password for the user account (should be hashed)
	// @json password
	// @example "securepassword123"
	// @swaggerignore
	Password string `json:"password" swaggerignore:"true"`
}

func (u *UserSchema) Create(user *UserSchema) *UserSchema {
	DB.Create(&user)
	return user
}

func (u *UserSchema) GetByEmail(email string) *UserSchema {
	var user UserSchema
	DB.Where("email = ?", email).First(&user)
	return &user
}

func (u *UserSchema) GetByID(id uint) *UserSchema {
	var user UserSchema
	DB.First(&user, id)
	return &user
}

var User = &UserSchema{}
