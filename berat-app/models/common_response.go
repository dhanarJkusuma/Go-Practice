package models

type CommonResponse struct {
	IsError bool
	Message string
	Data    interface{}
}
