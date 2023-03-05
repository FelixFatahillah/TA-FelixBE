package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"os"
	"product-api/controller"
	"product-api/db"
	"product-api/middleware"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	db.SetupDatabaseConnection()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "product-api",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		SendCookie:      true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// r.Use(CORSMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001", "http://localhost:3000"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	r.POST("/register", controller.Register)
	r.POST("/login", authMiddleware.LoginHandler)
	admin := r.Group("/admin", controller.LoginAdmin)
	admin.POST("/login", authMiddleware.LoginHandler)
	r.GET("/logout", authMiddleware.LogoutHandler)
	r.GET("/tours", controller.GetAllDestinations)
	r.GET("/picks", controller.GetAllPicks)
	r.GET("/cities", controller.GetAllCities)
	r.GET("/pick", controller.GetPickById)
	r.GET("/tour", controller.GetDestinationByPlace)
	r.POST("/order/package", controller.CustomerCreateOrder)
	r.POST("/order/custom", controller.CustomerCreateOrderCustom)
	r.GET("/transportations-custom", controller.GetAllTransportationsCustom)
	r.GET("/transportation-custom", controller.GetTransportationByIDCustom)
	r.POST("/upload", controller.FileUpload())
	r.PATCH("/upload-payment/patch", controller.UpdatePaymentByID)

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/self/:username", controller.GetUserByUsername)

		auth.POST("/destination", controller.PostDestination)
		auth.GET("/destinations", controller.GetAllDestinations)
		auth.GET("/destination", controller.GetDestinationByID)
		auth.PATCH("/destination/patch", controller.UpdateDestinationByID)
		auth.DELETE("/destination/delete", controller.DeleteDestinationByID)

		auth.POST("/user", controller.PostUser)
		auth.GET("/users", controller.GetAllUsers)
		auth.GET("/user", controller.GetUserByID)
		auth.PATCH("/user/patch", controller.UpdateUserByID)
		auth.DELETE("/user/delete", controller.DeleteUserByID)

		auth.POST("/order", controller.PostOrder)
		auth.GET("/orders", controller.GetAllOrders)
		auth.GET("/order", controller.GetOrderByID)
		auth.PATCH("/order/patch", controller.UpdateOrderByID)
		auth.DELETE("/order/delete", controller.DeleteOrderByID)

		auth.POST("/order_item", controller.PostOrderItem)
		auth.GET("/order_items", controller.GetAllOrderItems)
		auth.GET("/order_item", controller.GetOrderItemByID)
		auth.PATCH("/order_item/patch", controller.UpdateOrderItemByID)
		auth.DELETE("/order_item/delete", controller.DeleteOrderItemByID)

		auth.POST("/product_package", controller.PostProductPackage)
		auth.GET("/product_packages", controller.GetAllProductPackages)
		auth.GET("/product_package", controller.GetProductPackageByID)
		auth.PATCH("/product_package/patch", controller.UpdateProductPackageByID)
		auth.DELETE("/product_package/delete", controller.DeleteProductPackageByID)

		auth.POST("/transportation", controller.PostTransportation)
		auth.GET("/transportations", controller.GetAllTransportations)
		auth.GET("/transportation", controller.GetTransportationByID)
		auth.PATCH("/transportation/patch", controller.UpdateTransportationByID)
		auth.DELETE("/transportation/delete", controller.DeleteTransportationByID)

	}

	log.Fatal(r.Run(":" + port))
}
