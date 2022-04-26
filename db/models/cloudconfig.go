package models

import (
	"github.com/tnyeanderson/pixie/types"
	"gorm.io/gorm"
)

type CloudConfig struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Path         string `gorm:"unique;not null"`
	LastAccessed types.NullTime
}
