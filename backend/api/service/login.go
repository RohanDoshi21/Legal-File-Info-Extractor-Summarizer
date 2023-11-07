package service

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	M "github.com/word-extractor/word-extractor-apis/my_models"
	T "github.com/word-extractor/word-extractor-apis/types"
	U "github.com/word-extractor/word-extractor-apis/util"
)

type UserBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type User struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func Login(user *UserBody, trx *sql.Tx) (*M.User, *T.ServiceError) {

	dbCtx := context.Background()
	// Check if user exists
	exists, err := M.Users(
		M.UserWhere.Email.EQ(user.Email),
	).Exists(dbCtx, trx)

	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user",
		}
	}

	if !exists {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrUnauthorized.Code,
			Message: "User does not exist",
		}
	}
	// Get the user details
	userDetails, err := M.Users(
		M.UserWhere.Email.EQ(user.Email),
	).One(dbCtx, trx)

	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user",
		}
	}

	// Compare the password
	isValidPassword := U.ComparePassword(userDetails.Password, user.Password)
	
	if !isValidPassword {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrUnauthorized.Code,
			Message: "Invalid password",
		}
	}

	return userDetails, nil

}
