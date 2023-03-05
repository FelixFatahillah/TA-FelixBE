package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Email       string
	Phone       string
	Fullname    string
}
