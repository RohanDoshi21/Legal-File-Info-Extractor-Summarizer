package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	S "github.com/word-extractor/word-extractor-apis/api/service"
	C "github.com/word-extractor/word-extractor-apis/config"
	H "github.com/word-extractor/word-extractor-apis/handler"
	U "github.com/word-extractor/word-extractor-apis/util"
)

func Login(ctx *fiber.Ctx) error {
	// Get the body from the context
	userBody := ctx.Locals("body").(*S.UserBody)

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	authenticatedUser, err := S.Login(userBody, pgTrx)

	if err != nil {
		return H.BuildError(ctx, err.Message, err.Code, err.Error)
	}
	// Add claims to the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       authenticatedUser.ID,
		"email":    authenticatedUser.Email,
		"is_admin": authenticatedUser.Isadmin,
	})
	// Generate the token string
	tokenString, signingErr := token.SignedString([]byte(C.Conf.JwtSecret))

	if signingErr != nil {
		return H.BuildError(ctx, "Error generating token", fiber.StatusInternalServerError, signingErr)
	}

	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"token": tokenString,
	})

}
