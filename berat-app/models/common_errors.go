package models

import "errors"

var (
	ERR_NOT_FOUND   = errors.New("Record not found")
	ERR_INTERNAL_DB = errors.New("Database internal error")
)
