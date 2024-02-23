package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateSpot(c *gin.Context) {
	var spot models.Spot
	tripid := c.Param("tripid")
	if err := c.ShouldBindJSON(&spot); err != nil {
		log.Fatal("Failed to bind json")
	}
	spot.TripID = tripid
	
	if err := spot.CreateSpot(); err != nil {
		log.Fatal("Failed to create new spot")
	}

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

func UpdateSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")
	spot.FindSpotByID(id)

	if err := c.ShouldBindJSON(&spot); err != nil {
		log.Fatal("Failed to bind json")
	}
	if err := spot.UpdateSpotByID(); err != nil {
		log.Fatal("Failed to update user")
	}

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