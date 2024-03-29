package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/exceptions"
)

func FindPrefectureByID(c *gin.Context) {
	var prefecture models.Prefecture
	id := c.Param("prefectureid")
	prefecture.FindPrefectureByID(id)
	if err := prefecture.FindPrefectureByID(id); err == exceptions.NotFound {
		c.String(http.StatusNotFound, "Not Found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"prefecture": prefecture,
	})
}

func GetAllPrefectures(c *gin.Context) {
	var prefectures models.Prefectures
	if err := prefectures.GetAllPrefectures(); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
        return
	}
	c.JSON(http.StatusOK, gin.H{
		"prefectures": prefectures,
	})
}