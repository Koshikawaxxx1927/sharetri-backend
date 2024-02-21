package controllers

import (
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateSpot(c *gin.Context) {
	tripid := c.Param("tripid")
	spot := models.Spot{
		TripID: tripid,
		Name: "千葉",
		Date: time.Now(),
		StartTime: time.Now(),
		EndTime: time.Now(),
		Cost: 1000,
		Memo: "楽しかった",
	}
	
	if err := spot.CreateSpot(); err != nil {
		log.Fatal("Failed to create new spot")
	}
	
	// config.GetDB().First(&spot, 3)

	c.JSON(http.StatusOK, gin.H{
		"code": spot,
	})
}

func FindSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")
	spot.FindSpotByID(id)

	c.JSON(http.StatusOK, gin.H{
		"code": spot,
	})
}

func DeleteSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")

	if err := spot.DeleteSpotByID(id); err != nil {
		log.Fatal("Failed to delete the spot")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": spot,
	})
}