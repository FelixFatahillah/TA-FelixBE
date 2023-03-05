package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Fullname             string
	Status               string
	TransportationID     int
	TransportationQty    int
	TotalPrice           int
	DestinationPackageID *int
	IsPackage            bool
	Email                string
	Phone                string
	OrderDate            *time.Time
	PictureUrl           string
	Duration             int
}
