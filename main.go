package main

import (
	"log"
	"sync"

	"github.com/example/intern/controllers"
	"github.com/example/intern/database"
	"github.com/example/intern/middleware"
	"github.com/example/intern/models"
	"github.com/example/intern/services"
	"github.com/example/intern/utils"
	"github.com/gin-gonic/gin"
)

var productChannel = make(chan string)

func main() {
	go utils.LogServiceStatus()

	var wg sync.WaitGroup
	wg.Add(1)
	go utils.ProductMessageReceiver(productChannel, &wg)

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
	r.POST("/products", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		ProductController.CreateProduct(c, productChannel)
	})
	r.GET("/products", middleware.JWTAuthMiddleware(), ProductController.GetOwnProducts)
	r.DELETE("/products/:id", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		ProductController.DeleteProduct(c, productChannel)
	})
	r.PUT("/products/:id", middleware.JWTAuthMiddleware(), ProductController.UpdateProduct)

	port := "5000"
	log.Printf("Server is running on port %s", port)

	defer close(productChannel)

	r.Run(":" + port)
	wg.Wait()
}
