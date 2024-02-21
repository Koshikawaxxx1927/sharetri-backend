package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/http/controllers"
)

func Router() {
	router := gin.Default()
	router.Use(cors.Default())

	// router.GET("/", controllers.Index)
	// router.GET("/product", controllers.ProductIndex)
	router.GET("/createuser", controllers.CreateUser)
	router.GET("/finduser/:id", controllers.FindUserByID)
	router.GET("/deleteuser/:id", controllers.DeleteUserByID)
	router.GET("/findprefecture/:id", controllers.FindPrefectureByID)

	router.GET("/addtrip/:userid", controllers.CreateTrip)
	router.GET("/findtrip/:id", controllers.FindTripByID)
	router.GET("/deletetrip/:id", controllers.DeleteTripByID)

	router.GET("/addspot/:tripid", controllers.CreateSpot)
	router.GET("/findspot/:id", controllers.FindSpotByID)
	router.GET("/deletespot/:id", controllers.DeleteSpotByID)
	

	router.Run(":8080")
}