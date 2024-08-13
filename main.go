package main

import (
	"log"

	"github.com/example/intern/controllers"
	"github.com/example/intern/database"
	"github.com/example/intern/middleware"
	"github.com/example/intern/models"
	"github.com/example/intern/services"
	"github.com/example/intern/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	go utils.LogServiceStatus()

	admin := models.UserModel{
		Email:    "admin@example.com",
		Password: "securepassword",
	}

	database.InitDB()
	db := database.GetDB()
	db.AutoMigrate(&models.UserModel{}, &models.ProductModel{})

	utils.SeedAdminUser(db, &admin)

	r := gin.Default()

	/// user
	userService := services.NewUserService(db)
	UserController := controllers.NewUserController(userService)
	r.POST("/sign-up", UserController.CreateUser)
	r.POST("/sign-in", UserController.SignInUser)

	/// product
	productService := services.NewProductService(db)
	ProductController := controllers.NewProductController(productService)
	r.POST("/products", middleware.JWTAuthMiddleware(), ProductController.CreateProduct)
	r.GET("/products", middleware.JWTAuthMiddleware(), ProductController.GetOwnProducts)
	// r.GET("/products/:id", middleware.JWTAuthMiddleware(), getProduct)
	// r.PUT("/products/:id", middleware.JWTAuthMiddleware(), updateProduct)
	// r.DELETE("/products/:id", middleware.JWTAuthMiddleware(), deleteProduct)

	port := "5000"
	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}
