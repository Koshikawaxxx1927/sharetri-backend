package routes

import (
	// "time"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/http/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()
	// CORSミドルウェアを追加
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // 許可するオリジンを設定
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 許可するHTTPメソッドを設定
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	// For users
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userid", controllers.FindUserByID)
	router.DELETE("/user/:userid", controllers.DeleteUserByID)
	router.PUT("/user/:userid", controllers.UpdateUserByID)

	// router.POST("/usericon/:userid", controllers.UploadUserIcon)
	// router.DELETE("/usericon/:userid", controllers.DeleteUserIcon)

	// For prefectures
	router.GET("/prefecture/:prefectureid", controllers.FindPrefectureByID)
	router.GET("/prefecturelist", controllers.GetAllPrefectures)

	// For trips
	router.POST("/trip/:userid", controllers.CreateTrip)
	router.PUT("/trip/:tripid", controllers.UpdateTripByID)
	router.GET("/trip/:tripid", controllers.FindTripByID)
	router.DELETE("/trip/:tripid", controllers.DeleteTripByID)

	router.POST("/tripimage/:tripid", controllers.UploadTripImage)
	router.DELETE("/tripimage/:tripid", controllers.DeleteTripImage)
	router.GET("/tripimage/:tripid", controllers.GetTripImage)

	router.GET("/tripalllist", controllers.GetAllTrips)
	router.GET("/triplist", controllers.GetTrips)

	// For spots
	router.POST("/spot/:tripid", controllers.CreateSpot)
	router.GET("/spot/:spotid", controllers.FindSpotByID)
	router.PUT("/spot/:spotid", controllers.UpdateSpotByID)
	router.DELETE("/spot/:spotid", controllers.DeleteSpotByID)

	router.GET("/spotlist/:tripid", controllers.GetSpotsListByTripID)

	router.POST("/spotimage/:spotid", controllers.UploadSpotImage)
	router.DELETE("/spotimage/:spotid", controllers.DeleteSpotImage)
	router.GET("/spotimage/:spotid", controllers.GetSpotImage)
	
	return router
}