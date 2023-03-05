package form

import "gorm.io/gorm"

type Transportation struct {
	gorm.Model
	Type  string `json:"Type" binding:"required"`
	Size  int    `json:"Size" binding:"required"`
	Price int    `json:"Price" binding:"required"`
}
