package controllers

import (
	"log"
	// "time"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func CreateSpot(c *gin.Context) {
	tripid := c.Param("tripid")

	name := c.PostForm("name")
	date := c.PostForm("date")
	starttime := c.PostForm("starttime")
	endtime := c.PostForm("endtime")
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	memo := c.PostForm("memo")

	spot := models.Spot{
		TripID: tripid,
		Name: name,
		Date: utils.StringToTime(date),
		StartTime: utils.StringToTime(starttime),
		EndTime: utils.StringToTime(endtime),
		Cost: cost,
		Memo: memo,
	}
	
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

	tripid := c.PostForm("tripid")
	name := c.PostForm("name")
	date := c.PostForm("date")
	starttime := c.PostForm("starttime")
	endtime := c.PostForm("endtime")
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	memo := c.PostForm("memo")
	
	
	spot.TripID = tripid
	spot.Name = name
	spot.Date = utils.StringToTime(date)
	spot.StartTime = utils.StringToTime(starttime)
	spot.EndTime = utils.StringToTime(endtime)
	spot.Cost = cost
	spot.Memo = memo

	spot.UpdateSpotByID()

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