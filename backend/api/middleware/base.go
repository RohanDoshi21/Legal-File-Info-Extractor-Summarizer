package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	C "github.com/word-extractor/word-extractor-apis/config"
	DB "github.com/word-extractor/word-extractor-apis/db"
	H "github.com/word-extractor/word-extractor-apis/handler"
	T "github.com/word-extractor/word-extractor-apis/types"
	U "github.com/word-extractor/word-extractor-apis/util"
)

type Claims struct {
	jwt.RegisteredClaims
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Usertype bool   `json:"is_admin"`
}

func AuthenticateUser(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()

	authHeader := headers["Authorization"]

	if authHeader == "" {
		msg := "Authorization header is missing!"
		return H.BuildError(ctx, msg, fiber.StatusBadRequest, nil)
	}

	splitToken := strings.Split(authHeader, " ")

	if len(splitToken) != 2 {
		msg := "Token is missing in the header!"
		return H.BuildError(ctx, msg, fiber.StatusBadRequest, nil)
	}

	authKind := splitToken[0]

	if authKind != "Bearer" {
		msg := "Invalid authorization scheme!"
		return H.BuildError(ctx, msg, fiber.StatusBadRequest, nil)
	}

	token := splitToken[1]
	user, authErr := HandleBearerAuth(token)
	if authErr != nil {
		return H.BuildError(ctx, "Invalid or expired token/key!", fiber.ErrUnauthorized.Code, authErr)
	}
	ctx.Locals("user", user)
	ctx.Next()
	return nil

}

func AuthenticateAdmin(ctx *fiber.Ctx) error {
	user := U.GetAuthUser(ctx)
	if !user.IsAdmin {
		return H.BuildError(ctx, "Unauthorized access!", fiber.ErrForbidden.Code, nil)
	}
	ctx.Next()
	return nil
}

func HandleBearerAuth(token string) (*T.AuthUser, error) {
	user := &Claims{}
	byte_token, err := jwt.ParseWithClaims(token, user, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(C.Conf.JwtSecret), nil
	},
	)
	if byte_token != nil {
		if claims, ok := byte_token.Claims.(*Claims); ok && byte_token.Valid {
			
			authUser := T.AuthUser{
				User: T.User{
					Id:      claims.Id,
					Email:   claims.Email,
					IsAdmin: claims.Usertype,
				},
			}

			return &authUser, nil
		}
	}
	return nil, err
}

func Transaction(ctx *fiber.Ctx) error {
	dbCtx := context.Background()
	pgTrx, err := DB.PGTransaction(dbCtx)

	if err != nil {
		return H.BuildError(ctx, "Failed to initiate a transaction!", fiber.ErrInternalServerError.Code, err)
	}

	ctx.Locals("pgTrx", pgTrx)

	ctx.Next()

	return nil
}
