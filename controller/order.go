package controller

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/helper"
	"product-api/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var product []model.Order
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostOrder(c *gin.Context) {
	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Order{
		Fullname:             input.Fullname,
		Status:               input.Status,
		TransportationID:     input.TransportationID,
		TransportationQty:    input.TransportationQty,
		TotalPrice:           input.TotalPrice,
		DestinationPackageID: input.DestinationPackageID,
		IsPackage:            input.IsPackage,
		Email:                input.Email,
		Phone:                input.Phone,
		OrderDate:            input.OrderDate,
		Duration:             input.Duration,
		PictureUrl:           input.PictureUrl,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productInput := model.Order{
		Model:                gorm.Model{
			ID:        input.ID,
		},
		Status:               input.Status,

	}

	fmt.Println("tes aja", input)
	fmt.Println("tes aja", input)
	db.DB.Model(&product).Updates(productInput)

	helper.SentEmailConfirmation(product.Fullname, product.OrderDate.Format("2006-01-02 15:04:05"), product.Email)

	c.JSON(http.StatusOK, product)
}

func DeleteOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func CustomerCreateOrder(c *gin.Context) {
	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Order{
		Fullname:             input.Fullname,
		Status:               input.Status,
		TransportationID:     input.TransportationID,
		TransportationQty:    input.TransportationQty,
		TotalPrice:           input.TotalPrice,
		DestinationPackageID: input.DestinationPackageID,
		IsPackage:            input.IsPackage,
		Email:                input.Email,
		Phone:                input.Phone,
		OrderDate:            input.OrderDate,
		Duration:             input.Duration,
	}
	db.DB.Create(&product)

	strID := strconv.Itoa(int(product.ID))

	helper.SentEmail(strID, product.Email)

	customerData := form.Customer{
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}

	err := CreateCustomer(customerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "error create customer"})

	}

	c.JSON(http.StatusCreated, product)
}

func CustomerCreateOrderCustom(c *gin.Context) {
	type Result struct {
		Status               string     `json:"Status" binding:"required"`
		TotalPrice           int        `json:"TotalPrice" binding:"required"`
		TransportationID     int        `json:"TransportationID" binding:"required"`
		TransportationQty    int        `json:"TransportationQty" binding:"required"`
		Fullname             string     `json:"Fullname"`
		DestinationPackageID *int       `json:"DestinationPackageID"`
		IsPackage            bool       `json:"IsPackage"`
		Email                string     `json:"Email"`
		Phone                string     `json:"Phone"`
		OrderDate            *time.Time `json:"OrderDate"`
		Duration             int        `json:"Duration"`
		ArrOrderItem         []int      `json:"ArrOrderItem"`
	}

	var result Result

	if err := c.BindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Order{
		Fullname:             result.Fullname,
		Status:               result.Status,
		TransportationID:     result.TransportationID,
		TransportationQty:    result.TransportationQty,
		TotalPrice:           result.TotalPrice,
		DestinationPackageID: result.DestinationPackageID,
		IsPackage:            result.IsPackage,
		Email:                result.Email,
		Phone:                result.Phone,
		OrderDate:            result.OrderDate,
		Duration:             result.Duration,
	}
	db.DB.Create(&product)

	fmt.Println(product.ID)

	for i := range  result.ArrOrderItem {
		productItem := model.OrderItem{
			OrderID:       int(product.ID),
			DestinationID: result.ArrOrderItem[i],
		}
		db.DB.Create(&productItem)
	}

	strID := strconv.Itoa(int(product.ID))

	helper.SentEmail(strID, product.Email)

	customerData := form.Customer{
		Fullname: result.Fullname,
		Email:    result.Email,
		Phone:    result.Phone,
	}

	err := CreateCustomer(customerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "error create customer"})

	}

	c.JSON(http.StatusCreated, product)
}

func UpdatePaymentByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("tes aja", input)
	fmt.Println("tes aja", input.PictureUrl)
	fmt.Println("tes aja", input.ID)

	updateInput := model.Order{
		Model:                gorm.Model{
			ID:        input.ID,
		},
		PictureUrl:           input.PictureUrl,
	}

	db.DB.Model(&product).Updates(updateInput)

	c.JSON(http.StatusOK, product)
}

func CreateCustomer( input form.Customer )(err error) {


	product := model.Customer{
		Email:    input.Email,
		Phone:    input.Phone,
		Fullname: input.Fullname,
	}
	err = db.DB.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}
