package ongkir

type OngkirUsecase interface {
	Calculate(ongkirType int, weight float64) float64
}
