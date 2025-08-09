package models

import (
	"gorm.io/gorm"
)

type Snippet struct {
	gorm.Model
	Text string
}
