package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/word-extractor/word-extractor-apis/api/routes"
)

func InjectRoutes(app *fiber.App) {
	base := app.Group("/")
	routes.SetupRoutes(base)
}
