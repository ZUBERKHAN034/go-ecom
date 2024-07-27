package models

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/db"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

type UserSchema struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func init() {
	db.Connect()
	dbInstance = db.GetDB()
	dbInstance.AutoMigrate(&UserSchema{})
}

func (u *UserSchema) Create(user *UserSchema) *UserSchema {
	dbInstance.Create(&user)
	return user
}

func (u *UserSchema) GetByEmail(email string) *UserSchema {
	var user UserSchema
	dbInstance.Where("email = ?", email).First(&user)
	return &user
}

func (u *UserSchema) GetByID(id uint) *UserSchema {
	var user UserSchema
	dbInstance.First(&user, id)
	return &user
}

var User = &UserSchema{}
