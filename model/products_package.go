package model

import (
	"gorm.io/gorm"
)

type ProductPackage struct {
	gorm.Model
	Package          string
	DestinationCity  string
	PricePackage     int
	Description      string
	TransportationID int
	PictureUrl       string
	Duration         int
}
