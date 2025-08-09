package models

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Text string
}
