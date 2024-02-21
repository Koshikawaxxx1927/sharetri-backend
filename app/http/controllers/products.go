package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/Koshikawaxxx1927/sharetri-backend/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func ProductIndex(c *gin.Context) {
	product := models.Product{
		Code: "D42", Price: 100,
	}
	
	if err := product.Create(); err != nil {
		log.Fatal("Failed to create new product")
	}
	
	config.GetDB().First(&product, 1)

	c.JSON(http.StatusOK, gin.H{
		"code": product,
	})
}