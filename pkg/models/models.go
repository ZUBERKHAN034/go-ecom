package models

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/db"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func init() {
	DBInstance, err := db.Connect()
	if err != nil {
		panic("failed to connect to database [ REASON ] => " + err.Error())
	}

	// Migrate the schemas
	DBInstance.AutoMigrate(&UserSchema{})
	DBInstance.AutoMigrate(&ProductSchema{})
}
