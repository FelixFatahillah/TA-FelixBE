package form

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Place       string `json:"Place" binding:"required"`
	PlaceOption string `json:"PlaceOption" binding:"required"`
	Price       int    `json:"Price" binding:"required"`
	PictureUrl  string `json:"PictureUrl" binding:"required"`
	Description string `json:"Description" binding:"required"`
}
