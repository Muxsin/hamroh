package models

import (
	"gorm.io/gorm"
)

type Tutorial struct {
	gorm.Model
	Text string
}
