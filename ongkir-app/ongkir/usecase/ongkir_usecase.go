package usecase

import (
	"sirclo-test/ongkir-app/models"
	OngkirModule "sirclo-test/ongkir-app/ongkir"
)

type ongkirUsecase struct {
}

func NewOngkirUsecase() OngkirModule.OngkirUsecase {
	return &ongkirUsecase{}
}

func (ou *ongkirUsecase) Calculate(ongkirType int, weight float64) float64 {
	var result float64
	switch ongkirType {
	case models.TYPE_REGULAR:
		result = weight * 9000
	case models.TYPE_EXPRESS:
		result = weight * 9000 * 2
	case models.TYPE_INSTANT:
		result = weight * 9000 * 5
	}
	return result
}
