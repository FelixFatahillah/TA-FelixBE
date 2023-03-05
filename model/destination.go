package model

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Place       string
	PlaceOption string
	Price       int
	PictureUrl  string
	Description string
}
