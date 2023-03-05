package form

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID       int `json:"OrderID" binding:"required"`
	DestinationID int   `json:"DestinationID" binding:"required"`
}
