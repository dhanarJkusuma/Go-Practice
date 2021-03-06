package weight

import "sirclo-test/berat-app/models"

type WeightRepository interface {
	CreateRecord(data *models.Weight) (*models.Weight, error)
	ListRecord(page int, size int) ([]models.Weight, error)
	CountRecord() (int, error)
	DetailRecord(id int64) (*models.Weight, error)
	UpdateRecord(id int64, data models.Weight) (*models.Weight, error)
	DeleteRecord(id int64) error
}
