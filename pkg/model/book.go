package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	ISBN        string
	Author      string
}
