package controllers

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func FindPrefectureByID(c *gin.Context) {
	var prefecture models.Prefecture
	id := c.Param("id")
	prefecture.FindPrefectureByID(id)
	fmt.Println(prefecture)
	c.JSON(http.StatusOK, gin.H{
		"code": prefecture,
	})
}