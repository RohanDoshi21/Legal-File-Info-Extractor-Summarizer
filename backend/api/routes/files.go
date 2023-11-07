package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/word-extractor/word-extractor-apis/api/controller"
	mw "github.com/word-extractor/word-extractor-apis/api/middleware"
)

func SetupFileRoutes(router fiber.Router) {
	router.Post("/",mw.Transaction, controller.UploadFile)
	router.Get("/",mw.Transaction, controller.GetFiles)
}
