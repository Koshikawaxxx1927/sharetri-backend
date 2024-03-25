package controllers_test

import (
	"flag"
	"os"
	"testing"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/routes"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/database/seed"
)

func TestMain(m *testing.M) {
	env := flag.String("e", "development", "")
    flag.Parse()
	config.InitDB(*env, true, models.Trip{}, models.User{}, models.Prefecture{}, models.Spot{})
	seed.Seed()

	ret := m.Run()
	os.Exit(ret)
}

var Router = routes.Router()