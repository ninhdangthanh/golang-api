package controllers

import (
	"net/http"

	"github.com/example/intern/models"
	"github.com/example/intern/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserService()

	if err := userService.CreateUser(&user); err != nil {
		appErr, ok := err.(*services.AppError)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
			return
		}
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusOK, user)
}
