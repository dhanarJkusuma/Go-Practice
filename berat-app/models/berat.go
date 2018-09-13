package models

import (
	"time"
)

type Weight struct {
	ID          int64     `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Date        time.Time `gorm:"not null" json:"date"`
	WeightMin   float64   `gorm:"not null" json:"weightMin" form:"weightMin"`
	WeightMax   float64   `gorm:"not null" json:"weightMax" form:"weightMax"`
	RequestDate string    `gorm:"-" form:"requestDate"`
}
