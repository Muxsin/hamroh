package models

import (
	"gorm.io/gorm"
)

type Grocery struct {
	gorm.Model
	Text string
}
