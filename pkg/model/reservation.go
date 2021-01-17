package model

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	// Book Book
	BookID          int
	TelephoneNumber string
	DNI             string
	ReturnedAt      *time.Time
}
