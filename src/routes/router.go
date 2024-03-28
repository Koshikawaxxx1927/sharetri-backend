package routes

import (
	// "time"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/http/controllers"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/http/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	// CORSミドルウェアを追加
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // 許可するオリジンを設定
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 許可するHTTPメソッドを設定
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// For users
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userid", controllers.FindUserByID)
	router.DELETE("/user/:userid", controllers.DeleteUserByID)
	router.PUT("/user/:userid", controllers.UpdateUserByID)

	router.POST("/usericon/:userid", controllers.UploadUserIcon)
	router.DELETE("/usericon/:userid", controllers.DeleteUserIcon)
	router.GET("/usericon/:userid", controllers.GetUserIcon)

	// For prefectures
	router.GET("/prefecture/:prefectureid", controllers.FindPrefectureByID)
	router.GET("/prefecturelist", controllers.GetAllPrefectures)

	triplogin := router.Group("/trip/login/api/v1")
	tripuser := router.Group("/trip/user/api/v1")
	// For trips
	triplogin.Use(middleware.AuthMiddleware())
	{
		triplogin.POST("/trip/:userid", controllers.CreateTrip)
		triplogin.PUT("/trip/:tripid", controllers.UpdateTripByID)
	}
	tripuser.Use(middleware.AuthUserTripMiddleware())
	{
		tripuser.PUT("/trip/:tripid", controllers.UpdateTripByID)
		tripuser.DELETE("/trip/:tripid", controllers.DeleteTripByID)
		tripuser.POST("/tripimage/:tripid", controllers.UploadTripImage)
		tripuser.DELETE("/tripimage/:tripid", controllers.DeleteTripImage)
	}
	router.GET("/trip/:tripid", controllers.FindTripByID)
	router.GET("/tripimage/:tripid", controllers.GetTripImage)
	router.GET("/tripalllist", controllers.GetAllTrips)
	router.GET("/triplist", controllers.GetTrips)

	spotlogin := router.Group("/spot/login/api/v1")
	spotuser := router.Group("/spot/user/api/v1")
	// For spots
	spotlogin.Use(middleware.AuthMiddleware())
	{
		spotlogin.POST("/spot/:tripid", controllers.CreateSpot)
	}
	spotuser.Use(middleware.AuthUserSpotMiddleware())
	{
		spotuser.PUT("/spot/:spotid", controllers.UpdateSpotByID)
		spotuser.DELETE("/spot/:spotid", controllers.DeleteSpotByID)
		spotuser.POST("/spotimage/:spotid", controllers.UploadSpotImage)
		spotuser.DELETE("/spotimage/:spotid", controllers.DeleteSpotImage)
	}
	router.GET("/spotlist/:tripid", controllers.GetSpotsListByTripID)
	router.GET("/spot/:spotid", controllers.FindSpotByID)
	router.GET("/spotimage/:spotid", controllers.GetSpotImage)

	favoritelogin := router.Group("/favorite/login/api/v1")
	favoritelogin.Use(middleware.AuthMiddleware())
	{
		favoritelogin.POST("/favorite", controllers.CreateFavorite)
		favoritelogin.DELETE("/favorite/:favoriteid", controllers.DeleteFavoriteByID)
	}
	// router.POST("/favorite", controllers.CreateFavorite)
	// router.DELETE("/favorite/:favoriteid", controllers.DeleteFavoriteByID)
	router.GET("/favorite/uid/:uid", controllers.FindFavoritesByUid)
	router.GET("/favorite/tripid/:tripid", controllers.FindFavoritesByTripId)
	
	return router
}