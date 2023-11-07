package service

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	M "github.com/word-extractor/word-extractor-apis/my_models"
	T "github.com/word-extractor/word-extractor-apis/types"
	U "github.com/word-extractor/word-extractor-apis/util"
)

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(userDetails *RegisterUser, trx *sql.Tx) (*M.User, *T.ServiceError) {
	dbCtx := context.Background()

	// Check if user already exists
	exist, err := M.Users(
		M.UserWhere.Email.EQ(userDetails.Email),
	).Exists(dbCtx, trx)

	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user",
		}
	}

	if exist {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrBadRequest.Code,
			Message: "User already exists",
		}
	}
	// Hash the user password
	hashedPassword, hashErr := U.HashPassword(userDetails.Password)
	if hashErr != nil {
		return nil, &T.ServiceError{
			Error:   hashErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while hashing password",
		}
	}

	// Insert the user into the database
	user := &M.User{
		Email:    userDetails.Email,
		Password: hashedPassword,
	}

	insertErr := user.Insert(dbCtx, trx, boil.Infer())
	if insertErr != nil {
		return nil, &T.ServiceError{
			Error:   insertErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while inserting user",
		}
	}
	return user, nil

}
