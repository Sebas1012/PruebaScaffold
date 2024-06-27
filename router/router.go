package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sebas1012/go-rest-api/handler"
)

func Routes(app *fiber.App) {
	// Group se usa para agrupar rutas que contengan un prefijo, en teste caso el prefijo seria /hello.
	index := app.Group("/hello", logger.New())
	index.Get("/", handler.Index)
}
