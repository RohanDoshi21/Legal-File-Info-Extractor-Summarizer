package controller

import (
	"github.com/gofiber/fiber/v2"
	S "github.com/word-extractor/word-extractor-apis/api/service"
	H "github.com/word-extractor/word-extractor-apis/handler"
	U "github.com/word-extractor/word-extractor-apis/util"
)

func Register(ctx *fiber.Ctx) error {
	// Get the signup body from the request body
	signupBody := ctx.Locals("body").(*S.RegisterUser)

	// Get the transaction from the fiber context
	pgTrx := U.GetPGTrxFromFiberCtx(ctx)
	
	signedInUser, err := S.Register(signupBody,pgTrx)
	
	if err != nil {
		return H.BuildError(ctx, err.Message, err.Code, err.Error)
	}
	
	return H.Success(ctx, fiber.Map{
		"ok":   1,
		"user": signedInUser,
	})
}
