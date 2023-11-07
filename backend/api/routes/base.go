package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/word-extractor/word-extractor-apis/api/controller"
	mw "github.com/word-extractor/word-extractor-apis/api/middleware"

	S "github.com/word-extractor/word-extractor-apis/api/service"
)

func SetupRoutes(router fiber.Router) {
	//Test Route
	router.Get("/tests", controller.Test)

	router.Post("/register", mw.Validate(&S.RegisterUser{}), mw.Transaction, controller.Register)
	router.Post("/login", mw.Validate(&S.UserBody{}), mw.Transaction, controller.Login)
	//Protected Routes. Only authenticated users can access these routes
	router.Use(mw.AuthenticateUser)

	SetupFileRoutes(router.Group("/files"))
		// Protected Admin Routes. Only authenticated admin users can access these routes
		router.Use(mw.AuthenticateAdmin)

	SetupAdminRoutes(router.Group("/admin"))

}
