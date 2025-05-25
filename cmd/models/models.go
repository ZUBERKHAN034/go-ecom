package models

import (
	"github.com/ZUBERKHAN034/go-ecom/cmd/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = db.Connect()
	if err != nil {
		panic("failed to connect to database [ REASON ] => " + err.Error())
	}

	// Migrate the schemas
	DB.AutoMigrate(&UserSchema{})
	DB.AutoMigrate(&ProductSchema{})
	DB.AutoMigrate(&OrderSchema{})
	DB.AutoMigrate(&OrderItemSchema{})
}
