package main

import (
	"github.com/gofiber/fiber"
	"github.com/vanessadanu/Finpro-Golang.git/database"
	"github.com/vanessadanu/Finpro-Golang.git/route"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
