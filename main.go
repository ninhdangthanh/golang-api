package main

import (
	"log"
	"sync"

	"github.com/example/intern/controllers"
	"github.com/example/intern/database"
	_ "github.com/example/intern/docs"
	"github.com/example/intern/middleware"
	"github.com/example/intern/models"
	"github.com/example/intern/services"
	"github.com/example/intern/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var productChannel = make(chan string)

// @title           Gin Service
// @version         1.0
// @description     A management service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   Santosh Kumar
// @contact.url    https://twitter.com/sntshk
// @contact.email  sntshkmr60@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host      localhost:8080
// @BasePath  /
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
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

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

	port := "8080"
	log.Printf("Server is running on port %s", port)

	defer close(productChannel)

	r.Run(":" + port)
	wg.Wait()
}
