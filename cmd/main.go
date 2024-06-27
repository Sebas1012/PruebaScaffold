package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebas1012/go-rest-api/router"
)

func main() {
	app := fiber.New()
	router.Routes(app)
	app.Listen(":3000")
}
