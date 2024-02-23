package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateTrip(c *gin.Context) {
	userid := c.Param("userid")
	var trip models.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		log.Fatal("Failed to bind json")
	}
	trip.UserID = userid
	if err := trip.CreateTrip(); err != nil {
		log.Fatal("Failed to create new trip")
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": trip,
	})
}

func FindTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	trip.FindTripByID(id)

	c.JSON(http.StatusOK, gin.H{
		"code": trip,
	})
}

func UpdateTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	trip.FindTripByID(id)

	if err := c.ShouldBindJSON(&trip); err != nil {
		log.Fatal("Failed to bind json")
	}
	if err := trip.UpdateTripByID(); err != nil {
		log.Fatal("Failed to update trip")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": trip,
	})
}

func DeleteTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")

	if err := trip.DeleteTripByID(id); err != nil {
		log.Fatal("Failed to delete the trip")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": trip,
	})
}