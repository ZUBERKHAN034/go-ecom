package models

import (
	"github.com/ZUBERKHAN034/go-ecom/pkg/db"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func init() {
	db.Connect()
	DBInstance = db.GetDB()

	// Migrate the schemas
	DBInstance.AutoMigrate(&UserSchema{})
	DBInstance.AutoMigrate(&BookSchema{})
}
