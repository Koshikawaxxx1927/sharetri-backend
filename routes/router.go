package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/http/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.FindUserByID)
	router.DELETE("/user/:id", controllers.DeleteUserByID)
	router.PUT("/user/:id", controllers.UpdateUserByID)

	router.POST("/usericon/:id", controllers.UploadUserIcon)
	router.DELETE("/usericon/:id", controllers.DeleteUserIcon)

	router.GET("/prefecture/:id", controllers.FindPrefectureByID)

	router.POST("/trip/:userid", controllers.CreateTrip)
	router.PUT("/trip/:id", controllers.UpdateTripByID)
	router.GET("/trip/:id", controllers.FindTripByID)
	router.DELETE("/trip/:id", controllers.DeleteTripByID)

	router.POST("/tripimage/:id", controllers.UploadTripImage)
	router.DELETE("/tripimage/:id", controllers.DeleteTripImage)

	router.POST("/spot/:tripid", controllers.CreateSpot)
	router.GET("/spot/:id", controllers.FindSpotByID)
	router.PUT("/spot/:id", controllers.UpdateSpotByID)
	router.DELETE("/spot/:id", controllers.DeleteSpotByID)

	router.POST("/spotimage/:id", controllers.UploadSpotImage)
	router.DELETE("/spotimage/:id", controllers.DeleteSpotImage)
	
	return router
}