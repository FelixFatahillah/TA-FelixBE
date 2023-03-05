package controller

import (
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func GetAllProductPackages(c *gin.Context) {
	var product []model.ProductPackage
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetProductPackageByID(c *gin.Context) {
	var product model.ProductPackage
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostProductPackage(c *gin.Context) {
	var input form.ProductPackage
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.ProductPackage{
		Package:          input.Package,
		DestinationCity:  input.DestinationCity,
		PricePackage:     input.PricePackage,
		Description:      input.Description,
		TransportationID: input.TransportationID,
		PictureUrl:       input.PictureUrl,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateProductPackageByID(c *gin.Context) {
	var product model.ProductPackage
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.ProductPackage
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteProductPackageByID(c *gin.Context) {
	var product model.ProductPackage
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
