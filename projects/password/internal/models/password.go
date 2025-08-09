package models

import (
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	Text string
}
