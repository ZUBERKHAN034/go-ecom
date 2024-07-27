package models

import (
	"fmt"
	"log"

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
	result := dbInstance.Create(&user)
	if result.Error != nil {
		log.Printf("Failed to create user: %v", result.Error)
	}
	log.Println("User", user)
	return user
}

func (u *UserSchema) GetByEmail(email string) *UserSchema {
	var user UserSchema
	result := dbInstance.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Printf("Failed to get user by email: %v", result.Error)
	}
	log.Println("User", &user)
	return &user
}

func (u *UserSchema) GetByID(id uint) *UserSchema {
	var user UserSchema
	result := dbInstance.First(&user, id)
	if result.Error != nil {
		log.Printf("Failed to get user by ID: %v", result.Error)
	}
	log.Println("User", user)
	return &user
}

func (u *UserSchema) ComparePassword(password string) bool {
	fmt.Println("DB User Password", u.Password, "User Password", password)
	return u.Password == password
}

var User = &UserSchema{}
