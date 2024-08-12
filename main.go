package main

import (
	"log"

	"github.com/example/intern/controllers"
	"github.com/example/intern/database"
	"github.com/example/intern/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedAdminUser(db *gorm.DB, admin *models.UserModel) {
	var existingUser models.UserModel

	if err := db.Where("email = ?", admin.Email).First(&existingUser).Error; err == nil {
		log.Println("Admin user already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("securepassword"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin.Password = string(hashedPassword)

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	} else {
		log.Println("Admin user created successfully")
	}
}

func main() {
	admin := models.UserModel{
		Email:    "admin@example.com",
		Password: "securepassword",
	}

	database.InitDB()
	db := database.GetDB()
	db.AutoMigrate(&models.UserModel{}, &models.SortModel{}, &models.ProductModel{})

	seedAdminUser(db, &admin)

	r := gin.Default()
	r.POST("/sign-up", controllers.CreateUser)
	r.POST("/sign-in", controllers.SignInUser)

	port := "5000"
	log.Printf("Server is running on port %s", port)
	r.Run(":" + port)
}
