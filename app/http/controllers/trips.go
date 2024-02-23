package controllers

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func CreateTrip(c *gin.Context) {
	userid := c.Param("userid")

	prefectureid := c.PostForm("prefectureid")
	title := c.PostForm("title")
	startdate := 	c.PostForm("startdate")
	enddate := c.PostForm("enddate")
	memo := c.PostForm("memo")
	imagepath := c.PostForm("imagepath")
	ispublic, _ := strconv.ParseBool(c.PostForm("ispublic"))

	trip := models.Trip{
		UserID: userid,
		PrefectureID: prefectureid,
		Title: title,
		StartDate: utils.StringToTime(startdate),
		EndDate: utils.StringToTime(enddate),
		Memo: memo,
		ImagePath: imagepath,
		IsPublic: ispublic,
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

func UpdateTripByID(c *gin.Context) {
	var trip models.Trip
	id := c.Param("id")
	trip.FindTripByID(id)

	userid := c.PostForm("userid")
	prefectureid := c.PostForm("prefectureid")
	title := c.PostForm("title")
	startdate := c.PostForm("startdate")
	enddate := c.PostForm("enddate")
	memo := c.PostForm("memo")
	imagepath := c.PostForm("imagepath")
	ispublic := c.PostForm("ispublic")
	
	
	trip.UserID = userid
	trip.PrefectureID = prefectureid
	trip.Title = title
	trip.StartDate = utils.StringToTime(startdate)
	trip.EndDate = utils.StringToTime(enddate)
	trip.Memo = memo
	trip.ImagePath = imagepath
	trip.IsPublic, _ = strconv.ParseBool(ispublic)
	trip.UpdateTripByID()

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