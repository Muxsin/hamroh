package models

import (
	"gorm.io/gorm"
	"time"
)

type People struct {
	gorm.Model
	Name      string
	Surname   string
	Birthdate time.Time
}
