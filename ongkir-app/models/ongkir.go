package models

var (
	TYPE_REGULAR = 0
	TYPE_EXPRESS = 1
	TYPE_INSTANT = 2
)

type OngkirType struct {
	OngkirName string
	OngkirType int
}
