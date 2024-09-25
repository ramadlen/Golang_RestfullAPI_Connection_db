package main

import (
	"golang-restfull-api/db"
	"golang-restfull-api/db/migration"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		Prefork: true,
	})
	//migration
	migration.MigrationHandler()
	//database
	db.DBInit()
	//routes
	// routes.RouteHandler(app)
	err:= app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}