package form

import "gorm.io/gorm"

type ProductPackage struct {
	gorm.Model
	Package          string `json:"Package" binding:"required"`
	DestinationCity  string `json:"DestinationCity" binding:"required"`
	PricePackage     int    `json:"PricePackage" binding:"required"`
	Description      string `json:"Description" binding:"required"`
	TransportationID int    `json:"TransportationID" binding:"required"`
	PictureUrl       string `json:"PictureUrl" binding:"required"`
	Duration         int    `json:"Duration"`
}
