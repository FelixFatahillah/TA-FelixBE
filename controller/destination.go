package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
)

func GetAllDestinations(c *gin.Context) {
	var product []model.Destination
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetAllPicks(c *gin.Context) {
	type Result struct {
		ID              int    `gorm:"column:id"`
		DestinationCity string `gorm:"column:destination_city"`
		Package         string `gorm:"column:package"`
		PricePackage    int    `gorm:"column:price_package"`
		Description     string `gorm:"column:description"`
		PictureUrl      string `gorm:"column:picture_url"`
		Type            string `gorm:"column:type"`
		Size            int    `gorm:"column:size"`
	}
	var result []Result
	db.DB.Raw("SELECT a.id, a.destination_city, a.package, a.price_package, a.description, a.picture_url, b.type, b.size " +
		"FROM travel.product_packages a " +
		"LEFT JOIN travel.transportations b " +
		"ON a.transportation_id = b.id " +
		"WHERE a.deleted_at IS NULL;").Scan(&result)

	c.JSON(http.StatusOK, result)
}

func GetAllCities(c *gin.Context) {
	type Result struct {
		Place string `gorm:"place"`
	}
	var result []Result
	db.DB.Raw("SELECT d.place from travel.destinations d GROUP BY d.place").Scan(&result)

	c.JSON(http.StatusOK, result)
}

func GetPickById(c *gin.Context) {
	type Result struct {
		ID              int    `gorm:"column:id"`
		DestinationCity string `gorm:"column:destination_city"`
		Package         string `gorm:"column:package"`
		PricePackage    int    `gorm:"column:price_package"`
		Description     string `gorm:"column:description"`
		PictureUrl      string `gorm:"column:picture_url"`
		Type            string `gorm:"column:type"`
		Size            int    `gorm:"column:size"`
		Duration        int    `gorm:"duration"`
		TransportID     int    `gorm:"transport_id"`
	}
	var result Result
	db.DB.Raw("SELECT a.id, a.duration, a.destination_city, a.package, a.price_package, a.description, a.picture_url, b.type, b.size, b.id AS transport_id "+
		"FROM travel.product_packages a "+
		"LEFT JOIN travel.transportations b "+
		"ON a.transportation_id = b.id "+
		"WHERE a.id = ? AND a.deleted_at IS NULL;", c.Query("id")).Scan(&result)

	c.JSON(http.StatusOK, result)
}

func GetDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetDestinationByPlace(c *gin.Context) {
	var product []model.Destination
	if err := db.DB.Where("place = ?", c.Query("place")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostDestination(c *gin.Context) {
	var input form.Destination
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := model.Destination{
		Place:       input.Place,
		PlaceOption: input.PlaceOption,
		Price:       input.Price,
		PictureUrl:  input.PictureUrl,
		Description: input.Description,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Destination
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
