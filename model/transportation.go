package model

import "gorm.io/gorm"

type Transportation struct {
	gorm.Model
	Type  string
	Size  int
	Price int
}
