package controllers

import (
	"net/http"

	"github.com/example/intern/models"
	"github.com/example/intern/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product models.ProductModel

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// ?
	// if strings.TrimSpace(product.Name) == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Product name must not be blank"})
	// 	return
	// }

	product.UserID = userID.(uint)

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.ProductService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (ctrl *ProductController) GetOwnProducts(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	products, err := ctrl.ProductService.GetOwnProducts(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
