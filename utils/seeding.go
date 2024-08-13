package utils

import (
	"log"

	"github.com/example/intern/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB, admin *models.UserModel) {
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
