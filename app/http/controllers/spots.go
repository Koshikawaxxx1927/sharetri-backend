package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
)

func CreateSpot(c *gin.Context) {
	var spot models.Spot
	tripid := c.Param("tripid")
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	spot.TripID = tripid
	
	if err := spot.CreateSpot(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"spot": spot,
	})
}

func FindSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")
	if err := spot.FindSpotByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"spot": spot,
	})
}

func UpdateSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")
	if err := spot.FindSpotByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := spot.UpdateSpotByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"spot": spot,
	})
}

func DeleteSpotByID(c *gin.Context) {
	var spot models.Spot
	id := c.Param("id")

	if err := spot.FindSpotByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := spot.DeleteSpotByID(id); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"spot": spot,
	})
}