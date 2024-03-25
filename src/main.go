package main

import (
	"flag"
	"log"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/routes"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/database/seed"
)

func main() {
	env := flag.String("e", "development", "")
    flag.Parse()

	config.InitDB(*env, true, models.Trip{}, models.User{}, models.Prefecture{}, models.Spot{})
	seed.Seed()
	router := routes.Router()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to build server")
	}
}