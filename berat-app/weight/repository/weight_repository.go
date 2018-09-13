package repository

import (
	"sirclo-test/berat-app/models"
	WeightModule "sirclo-test/berat-app/weight"

	"github.com/jinzhu/gorm"
)

type sqlWeightRepository struct {
	DB *gorm.DB
}

func NewWeightRepository(db *gorm.DB) WeightModule.WeightRepository {
	return &sqlWeightRepository{
		DB: db,
	}
}

func (r *sqlWeightRepository) CreateRecord(data *models.Weight) (*models.Weight, error) {
	if dbe := r.DB.Create(data).Error; dbe != nil {
		return nil, models.ERR_INTERNAL_DB
	}
	return data, nil
}

func (r *sqlWeightRepository) ListRecord() ([]models.Weight, error) {
	var results []models.Weight
	if dbe := r.DB.Order("date desc").Find(&results).Error; dbe != nil {
		return nil, models.ERR_INTERNAL_DB
	}
	return results, nil
}

func (r *sqlWeightRepository) DetailRecord(id int64) (*models.Weight, error) {
	var result models.Weight
	if dbe := r.DB.Where("id = ?", id).Find(&result).Error; dbe != nil {
		return nil, models.ERR_NOT_FOUND
	}
	return &result, nil
}

func (r *sqlWeightRepository) UpdateRecord(id int64, data models.Weight) (*models.Weight, error) {
	existRecord, err := r.DetailRecord(id)
	if err != nil {
		return nil, err
	}
	existRecord.Date = data.Date
	existRecord.WeightMax = data.WeightMax
	existRecord.WeightMin = data.WeightMin
	if dbe := r.DB.Save(existRecord).Error; dbe != nil {
		return nil, models.ERR_INTERNAL_DB
	}
	return existRecord, nil
}

func (r *sqlWeightRepository) DeleteRecord(id int64) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Weight{}).Error; err != nil {
		return models.ERR_NOT_FOUND
	}
	return nil
}
