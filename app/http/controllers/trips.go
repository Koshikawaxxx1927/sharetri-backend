package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
)

func CreateTrip(c *gin.Context) {
	userid := c.Param("userid")
	var trip models.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	trip.UserID = userid
	if err := trip.CreateTrip(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"trip": trip,
	})
}

func FindTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	if err := trip.FindTripByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"trip": trip,
	})
}

func UpdateTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")

	if err := trip.FindTripByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := trip.UpdateTripByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"trip": trip,
	})
}

func DeleteTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")

	if err := trip.FindTripByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := trip.DeleteTripByID(id); err != nil {
		log.Fatal("Failed to delete the trip")
	}

	c.JSON(http.StatusOK, gin.H{
		"trip": trip,
	})
}