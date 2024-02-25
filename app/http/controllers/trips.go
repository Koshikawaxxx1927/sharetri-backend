package controllers

import (
	// "log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func CreateTrip(c *gin.Context) {
	userid := c.Param("userid")
	var user models.User
	if err := user.FindUserByID(userid); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	var trip models.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	trip.UserID = userid
	trip.ImagePath = ""
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
	imagepath := trip.ImagePath
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	trip.ImagePath = imagepath
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
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusOK, gin.H{
		"trip": trip,
	})
}

func UploadTripImage(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	if err := trip.FindTripByID(id); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err := utils.DeleteFile(trip.ImagePath)
	if err != nil && trip.ImagePath != "" {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	var image utils.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	outputFile := "storage/trips/" + id
	savePath, err := utils.SaveDecodedImage(image.EncodedData, outputFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	trip.ImagePath = savePath
	if err := trip.UpdateTripByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"trip": trip,
	})
}

func DeleteTripImage(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	if err := trip.FindTripByID(id); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	
	if err := utils.DeleteFile(trip.ImagePath); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	trip.ImagePath = ""
	if err := trip.UpdateTripByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"trip": trip,
	})
}