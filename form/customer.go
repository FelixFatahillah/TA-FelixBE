package form

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Fullname string `json:"Fullname"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
}
