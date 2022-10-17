package main

import (
	"log"

	"github.com/Theodule007/go_fiber_api/pkg/books"
	"github.com/Theodule007/go_fiber_api/pkg/common/config"
	"github.com/Theodule007/go_fiber_api/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
