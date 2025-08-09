package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Text string
}
