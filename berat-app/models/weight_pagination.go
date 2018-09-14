package models

type WeightPagination struct {
	Page       int
	Size       int
	LastPage   int
	TotalCount int
	Content    []Weight
}
