package models

import (
	"gorm.io/gorm"
)

type Flashcard struct {
	gorm.Model
	Text string
}
