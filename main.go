package main

import (
	"fmt"

	"github.com/example/intern/models"
)

func main() {
	admin := models.UserModel{
		Email:    "admin@example.com",
		Password: "securepassword", // Make sure to hash passwords in a real application
		Gender:   models.Male,
		Role:     models.AdminRole,
		Account:  models.Active,
	}

	fmt.Println(admin)

	fmt.Println("Hello, World!")
}
