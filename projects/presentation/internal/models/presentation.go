package models

import (
	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	Text string
}
