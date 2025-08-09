package models

import (
	"gorm.io/gorm"
)

type Palette struct {
	gorm.Model
	Text string
}
