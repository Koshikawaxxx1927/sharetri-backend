package main

import (
	"flag"
	"github.com/Koshikawaxxx1927/sharetri-backend/routes"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/database/seed"
)

func main() {
	env := flag.String("e", "development", "")
    flag.Parse()

	config.InitDB(*env, true, models.Trip{}, models.User{}, models.Prefecture{}, models.Spot{})
	seed.Seed()
	routes.Router()
}