package models

import (
	"gorm.io/gorm"
)

type UserSchema struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *UserSchema) Create(user *UserSchema) *UserSchema {
	DBInstance.Create(&user)
	return user
}

func (u *UserSchema) GetByEmail(email string) *UserSchema {
	var user UserSchema
	DBInstance.Where("email = ?", email).First(&user)
	return &user
}

func (u *UserSchema) GetByID(id uint) *UserSchema {
	var user UserSchema
	DBInstance.First(&user, id)
	return &user
}

var User = &UserSchema{}
