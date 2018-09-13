package weight

import "sirclo-test/berat-app/models"

type WeightUsecase interface {
	CreateRecord(data *models.Weight) (bool, string)
	ListRecord() []models.Weight
	DetailRecord(id int64) (bool, *models.Weight)
	UpdateRecord(id int64, data models.Weight) (bool, string)
	DeleteRecord(id int64) (bool, string)
}
