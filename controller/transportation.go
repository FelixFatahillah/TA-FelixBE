package controller

import (
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func GetAllTransportations(c *gin.Context) {
	var product []model.Transportation
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetAllTransportationsCustom(c *gin.Context) {
	var product []model.Transportation
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetTransportationByID(c *gin.Context) {
	var product model.Transportation
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetTransportationByIDCustom(c *gin.Context) {
	var product model.Transportation
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostTransportation(c *gin.Context) {
	var input form.Transportation
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Transportation{
		Type:  input.Type,
		Size:  input.Size,
		Price: input.Price,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateTransportationByID(c *gin.Context) {
	var product model.Transportation
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Transportation
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteTransportationByID(c *gin.Context) {
	var product model.Transportation
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
