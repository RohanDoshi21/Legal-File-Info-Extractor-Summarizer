package middleware

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	H "github.com/word-extractor/word-extractor-apis/handler"
)

func Validate(body any) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if err := ctx.BodyParser(body); err != nil {
			msg := "Failed to parse the body!"
			return H.BuildError(ctx, msg, fiber.ErrBadRequest.Code, err)
		}
		fmt.Println(body)
		validate := validator.New()
		err := validate.Struct(body)
		if err != nil {
			return H.BuildError(ctx, err.Error(), fiber.ErrBadRequest.Code, err)
		} 

			ctx.Locals("body", body)
		
		ctx.Next()
		return nil

	}

}
