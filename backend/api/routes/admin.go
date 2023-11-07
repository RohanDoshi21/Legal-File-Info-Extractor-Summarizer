package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/word-extractor/word-extractor-apis/api/controller"
	mw "github.com/word-extractor/word-extractor-apis/api/middleware"
	S "github.com/word-extractor/word-extractor-apis/api/service"
)

func SetupAdminRoutes(router fiber.Router) {
	router.Get("/users", mw.Transaction, controller.GetUsers)
	router.Get("/userdocs", mw.Transaction, controller.GetUserDocs)
	router.Patch("/userdoc", mw.Validate(&S.DocBody{}), mw.Transaction, controller.UpdateUserDoc)
}
