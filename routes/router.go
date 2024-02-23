package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/http/controllers"
)

func Router() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.FindUserByID)
	router.DELETE("/user/:id", controllers.DeleteUserByID)
	router.PUT("/user/:id", controllers.UpdateUserByID)


	router.GET("/prefecture/:id", controllers.FindPrefectureByID)

	router.POST("/trip/:userid", controllers.CreateTrip)
	router.PUT("/trip/:userid", controllers.UpdateTripByID)
	router.GET("/trip/:id", controllers.FindTripByID)
	router.DELETE("/trip/:id", controllers.DeleteTripByID)

	router.POST("/spot/:tripid", controllers.CreateSpot)
	router.GET("/spot/:id", controllers.FindSpotByID)
	router.PUT("/spot/:id", controllers.UpdateSpotByID)
	router.DELETE("/spot/:id", controllers.DeleteSpotByID)
	

	router.Run(":8080")
}