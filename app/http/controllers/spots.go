package controllers

import (
	// "fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/exceptions"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func CreateSpot(c *gin.Context) {
	var spot models.Spot
	tripid := c.Param("tripid")
	var trip models.Trip
	if err := trip.FindTripByID(tripid); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	spot.TripID = tripid
	spot.ImagePath = ""
	if err := spot.CreateSpot(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"spot": spot,
	})
}

func FindSpotByID(c *gin.Context) {
	var spot models.Spot
	spotid := c.Param("spotid")
	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
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
	spotid := c.Param("spotid")
	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	imagepath := spot.ImagePath
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	spot.ImagePath = imagepath
	if err := spot.UpdateSpotByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"spot": spot,
	})
}

func DeleteSpotByID(c *gin.Context) {
	var spot models.Spot
	spotid := c.Param("spotid")

	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := spot.DeleteSpotByID(spotid); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"spot": spot,
	})
}

func UploadSpotImage(c *gin.Context) {
	var spot models.Spot
	spotid := c.Param("spotid")
	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err := utils.DeleteFile(spot.ImagePath)
	if err != nil && spot.ImagePath != "" {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	var image utils.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	outputFile := utils.ProjectRoot + "/storage/spots/" + spotid
	savePath, err := utils.SaveDecodedImage(image.EncodedData, outputFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	spot.ImagePath = savePath
	if err := spot.UpdateSpotByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"spot": spot,
	})
}

func DeleteSpotImage(c *gin.Context) {
	var spot models.Spot
	spotid := c.Param("spotid")
	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	
	if err := utils.DeleteFile(spot.ImagePath); err != nil {
		c.String(http.StatusNotFound, "Not Found")
        return
	}
	spot.ImagePath = ""
	if err := spot.UpdateSpotByID(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusCreated, gin.H{
		"spot": spot,
	})
}

func GetSpotsByTripID(c *gin.Context) {
	var spots models.Spots
	tripid := c.Param("tripid")
	if err := spots.GetSpotsByTripID(tripid); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"spots": spots,
	})
}

func GetAllSpots(c *gin.Context) {
	var spots models.Spots
	if err := spots.GetAllSpots(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"spots": spots,
	})
}

func GetSpotsListByTripID(c *gin.Context) {
	var spots models.Spots
	var tripid, offset, limit int
	var err error
	if offset, err = strconv.Atoi(c.Query("offset")); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if limit, err = strconv.Atoi(c.Query("limit")); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if tripid, err = strconv.Atoi(c.Param("tripid")); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// ページング処理を行う
	if err := spots.GetSpotsListByTripID(tripid, offset, limit); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"spots": spots,
	})
}

func GetSpotImage(c *gin.Context) {
	var spot models.Spot
	spotid := c.Param("spotid")
	if err := spot.FindSpotByID(spotid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if spot.ImagePath == "" {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	c.File(spot.ImagePath)
}

