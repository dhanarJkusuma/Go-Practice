package usecase

import (
	"sirclo-test/berat-app/models"
	WeightModule "sirclo-test/berat-app/weight"
)

type sqlWeightUsecase struct {
	Repo WeightModule.WeightRepository
}

// function like constructor
func NewWeightUsecase(Repo WeightModule.WeightRepository) WeightModule.WeightUsecase {
	return &sqlWeightUsecase{
		Repo,
	}
}

// Create Record
func (u *sqlWeightUsecase) CreateRecord(data *models.Weight) (bool, string) {
	_, err := u.Repo.CreateRecord(data)
	if err != nil {
		return false, "Terjadi kesalahan dalam penginputan data."
	}
	return true, "Berhasil menambahkan data."
}

// Fetch Record with Pagination
func (u *sqlWeightUsecase) ListRecord(page int, size int) *models.WeightPagination {
	// get count
	count, err := u.Repo.CountRecord()
	if err != nil {
		// just return empty array if error is occur
		return &models.WeightPagination{
			Page:       page,
			Size:       size,
			Content:    []models.Weight{},
			TotalCount: 0,
		}
	}
	// get record
	data, err := u.Repo.ListRecord(page, size)
	if err != nil {
		// just return empty array if error is occur
		return &models.WeightPagination{
			Page:       page,
			Size:       size,
			Content:    []models.Weight{},
			TotalCount: 0,
		}
	}
	return &models.WeightPagination{
		Page:       page,
		Size:       size,
		Content:    data,
		TotalCount: count,
	}
}

// Detail Record
func (u *sqlWeightUsecase) DetailRecord(id int64) (bool, *models.Weight) {
	data, err := u.Repo.DetailRecord(id)
	if err != nil {
		return false, nil
	}
	return true, data
}

// Update Record
func (u *sqlWeightUsecase) UpdateRecord(id int64, data models.Weight) (bool, string) {
	_, err := u.Repo.UpdateRecord(id, data)
	if err != nil {
		return false, "Terjadi kesalahan saat mengubah data."
	}
	return true, "Data berhasil diubah."
}

// Delete Record
func (u *sqlWeightUsecase) DeleteRecord(id int64) (bool, string) {
	err := u.Repo.DeleteRecord(id)
	if err != nil {
		return false, "Terjadi kesalahan saat menghapus data, data tidak ditemukan."
	}
	return true, "Berhasil menghapus data"
}
