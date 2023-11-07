package controller

import (
	"github.com/gofiber/fiber/v2"
	H "github.com/word-extractor/word-extractor-apis/handler"
)

func Test(ctx *fiber.Ctx) error {
	return H.Success(ctx, fiber.Map{
		"ok": 1,
	})
}
