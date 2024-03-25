package controllers

import (
	"io"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/exceptions"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/utils"
	"strconv"
	"github.com/google/uuid"
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
	trip.Uid = userid
	trip.ImagePath = ""
	if err := trip.CreateTrip(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"trip": trip,
	})
}

func FindTripByID(c *gin.Context) {
	var trip models.Trip
	tripid := c.Param("tripid")
	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
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
	tripid := c.Param("tripid")

	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
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

	c.JSON(http.StatusCreated, gin.H{
		"trip": trip,
	})
}

func DeleteTripByID(c *gin.Context) {
	var trip models.Trip
	tripid := c.Param("tripid")

	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := trip.DeleteTripByID(tripid); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}

	c.JSON(http.StatusCreated, gin.H{
		"trip": trip,
	})
}

func UploadTripImage(c *gin.Context) {
	var trip models.Trip
	tripid := c.Param("tripid")
	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err := utils.DeleteFile(trip.ImagePath)
	if err != nil && trip.ImagePath != "" {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	outputDir := utils.ProjectRoot + "/storage/spots/" + tripid
	outputFile := utils.ProjectRoot + "/storage/trips/" + tripid + "/"+ uuid.New().String()
	os.Mkdir(outputDir, 0777)
	out, err := os.Create(outputFile)
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	trip.ImagePath = outputFile
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
	tripid := c.Param("tripid")
	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	
	if err := utils.DeleteFile(trip.ImagePath); err != nil {
		c.String(http.StatusNotFound, "Not Found")
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

func GetAllTrips(c *gin.Context) {
	var trips models.Trips
	if err := trips.GetAllTrips(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"trips": trips,
	})
}

func GetTrips(c *gin.Context) {
	var trips models.Trips
	var offset, limit int
	var err error
	if offset, err = strconv.Atoi(c.Query("offset")); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if limit, err = strconv.Atoi(c.Query("limit")); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// ページング処理を行う
	if err := trips.GetTrips(offset, limit); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"trips": trips,
	})
}

func GetTripImage(c *gin.Context) {
	var trip models.Trip
	tripid := c.Param("tripid")
	if err := trip.FindTripByID(tripid); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if trip.ImagePath == "" {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	c.File(trip.ImagePath)
}