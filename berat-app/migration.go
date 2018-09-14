package main

import (
	"sirclo-test/berat-app/models"

	"github.com/jinzhu/gorm"
)

// AutoMigration is function to make database migration
func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Weight{})
}
