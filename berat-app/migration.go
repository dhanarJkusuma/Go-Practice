package main

import (
	"sirclo-test/berat-app/models"

	"github.com/jinzhu/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Weight{})
}
