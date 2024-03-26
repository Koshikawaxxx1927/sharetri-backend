package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
)

// ログインしているかを判定
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginAuth(c)
		c.Next()
	}
}

func AuthUserTripMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginAuth(c)
		tripid := c.Param("tripid")
		tripAuthMiddleware(c, tripid)
		c.Next()
	}
}

func AuthUserSpotMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginAuth(c)
		spotid := c.Param("spotid")
		spotAuthMiddleware(c, spotid)
		c.Next()
	}
}

// ログインしていてSpotのUIDが合っているかを判定
func spotAuthMiddleware(c *gin.Context, spotid string) {
	authorization := c.Request.Header.Get("Authorization")
	authUid := strings.Split(authorization, ";")[1]
	uid := strings.Replace(authUid, "UID ", "", 1)

	var spot models.Spot
	if err := spot.FindSpotByID(spotid); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	var trip models.Trip
	if err := trip.FindTripByID(spot.TripID); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	log.Printf("UID: %s %s", uid, trip.Uid)
	if uid != trip.Uid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "(Unauthorized)",
		})
		return
	}
	log.Printf("Vertifed UID: %s\n", uid)
}

// ログインしていてTripのUIDが合っているかを判定
func tripAuthMiddleware(c *gin.Context, tripid string) {
	authorization := c.Request.Header.Get("Authorization")
	authUid := strings.Split(authorization, ";")[1]
	uid := strings.Replace(authUid, "UID ", "", 1)

	var trip models.Trip
	if err := trip.FindTripByID(tripid); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	log.Printf("UID: %s %s", uid, trip.Uid)
	if uid != trip.Uid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "(Unauthorized)",
		})
		return
	}
	log.Printf("Vertifed UID: %s\n", uid)
}

func loginAuth(c *gin.Context) {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("err: %v\n", err)
		os.Exit(1)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("err: %v\n", err)
		os.Exit(1)
	}

	authorization := c.Request.Header.Get("Authorization")
	authHandler := strings.Split(authorization, ";")[0] 
	idToken := strings.Replace(authHandler, "Bearer ", "", 1)

	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": err.Error(),
		})
		return
	}
	log.Printf("Vertifed ID token: %v\n", token)
}