package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Summary string
	Detail  string
}
