package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	lgr "github.com/sirupsen/logrus"
	C "github.com/word-extractor/word-extractor-apis/config"
	"github.com/word-extractor/word-extractor-apis/db"
	"github.com/word-extractor/word-extractor-apis/handler"
)

func main() {
	godotenv.Load(".env")

	configValues, configErr := C.New()
	if configErr != nil {
		lgr.Fatalln("Failed to initialize app-wise configuration!", configErr)
	}

	err := db.Init()

	if err != nil {
		lgr.Fatalln("Error while setting up DB connections!", err)
	}

	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
		BodyLimit:    1024 * 1024 * 1024, // Max upload size: 4MB
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000 , http://localhost:3001",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length, Authorization, X-Api-Key",
	}))

	app.Use(requestid.New())

	InjectRoutes(app)

	port := configValues.Port
	lgr.Fatal(app.Listen(":" + port))

}
