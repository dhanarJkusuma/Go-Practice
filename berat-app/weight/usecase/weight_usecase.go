package usecase

import (
	"sirclo-test/berat-app/models"
	WeightModule "sirclo-test/berat-app/weight"
)

type sqlWeightUsecase struct {
	Repo WeightModule.WeightRepository
}

func NewWeightUsecase(Repo WeightModule.WeightRepository) WeightModule.WeightUsecase {
	return &sqlWeightUsecase{
		Repo,
	}
}

func (u *sqlWeightUsecase) CreateRecord(data *models.Weight) (bool, string) {
	_, err := u.Repo.CreateRecord(data)
	if err != nil {
		return false, "Terjadi kesalahan dalam penginputan data."
	}
	return true, "Berhasil menambahkan data."
}

func (u *sqlWeightUsecase) ListRecord() []models.Weight {
	data, err := u.Repo.ListRecord()
	if err != nil {
		return []models.Weight{}
	}
	return data
}

func (u *sqlWeightUsecase) DetailRecord(id int64) (bool, *models.Weight) {
	data, err := u.Repo.DetailRecord(id)
	if err != nil {
		return false, nil
	}
	return true, data
}

func (u *sqlWeightUsecase) UpdateRecord(id int64, data models.Weight) (bool, string) {
	_, err := u.Repo.UpdateRecord(id, data)
	if err != nil {
		return false, "Terjadi kesalahan saat mengubah data."
	}
	return true, "Data berhasil diubah."
}

func (u *sqlWeightUsecase) DeleteRecord(id int64) (bool, string) {
	err := u.Repo.DeleteRecord(id)
	if err != nil {
		return false, "Terjadi kesalahan saat menghapus data, data tidak ditemukan."
	}
	return true, "Berhasil menghapus data"
}
