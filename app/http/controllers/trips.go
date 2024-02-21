package controllers

import (
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

	// "github.com/Koshikawaxxx1927/sharetri-backend/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func CreateTrip(c *gin.Context) {
	userId := c.Param("userid")
	trip := models.Trip{
		UserID: userId,
		PrefectureID: "12",
		Title: "千葉旅行",
		StartDate: time.Now(),
		EndDate: time.Now(),
		Memo: "ここは楽しい",
		IsPublic: true,
	}
	
	if err := trip.CreateTrip(); err != nil {
		log.Fatal("Failed to create new trip")
	}
	
	// config.GetDB().First(&trip, 3)

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