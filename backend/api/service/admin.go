package service

import (
	"context"
	"database/sql"

	"github.com/go-jose/go-jose/v3/json"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	M "github.com/word-extractor/word-extractor-apis/my_models"
	T "github.com/word-extractor/word-extractor-apis/types"
)

type UserFilter struct {
	Id []int `json:"id"`
}
type UserDocFilter struct {
	UserID int `json:"user_id"`
}

type DocBody struct {
	DocID       int               `json:"doc_id"`
	FileContent map[string]string `json:"file_content" validate:"required"`
	EditedBy    int               `json:"edited_by"`
}

func GetUsers(userFilter *UserFilter, trx *sql.Tx) (*M.UserSlice, *T.ServiceError) {
	dbCtx := context.Background()
	userIds := make([]interface{}, len(userFilter.Id))
	for i, id := range userFilter.Id {
		userIds[i] = id
	}
	query := []qm.QueryMod{}
	if userFilter.Id != nil {
		query = append(query, qm.WhereIn("id in ?", userIds...))
	}
	users, err := M.Users(query...).All(dbCtx, trx)
	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching users",
		}
	}
	return &users, nil
}

func GetUserDocs(user *UserDocFilter, trx *sql.Tx) (*M.DocumentSlice, *T.ServiceError) {
	dbCtx := context.Background()
	userexists, userErr := M.Users(
		M.UserWhere.ID.EQ(user.UserID),
	).Exists(dbCtx, trx)
	if userErr != nil {
		return nil, &T.ServiceError{
			Error:   userErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user",
		}
	}
	if !userexists {
		return nil, &T.ServiceError{
			Error:   nil,
			Code:    fiber.ErrBadRequest.Code,
			Message: "User does not exist",
		}
	}

	userDocs, err := M.Documents(
		M.DocumentWhere.OwnerID.EQ(user.UserID),
	).All(dbCtx, trx)
	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user docs",
		}
	}
	return &userDocs, nil
}

func UpdateUserDoc(docBody *DocBody, trx *sql.Tx) (int, *T.ServiceError) {
	dbCtx := context.Background()

	docExists, err := M.Documents(
		M.DocumentWhere.ID.EQ(docBody.DocID),
	).Exists(dbCtx, trx)
	if err != nil {
		return 0, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user docs",
		}
	}
	if !docExists {
		return 0, &T.ServiceError{
			Error:   nil,
			Code:    fiber.ErrBadRequest.Code,
			Message: "Doc does not exist",
		}
	}

	doc, docErr := M.Documents(
		M.DocumentWhere.ID.EQ(docBody.DocID),
	).One(dbCtx, trx)

	if docErr != nil {
		return 0, &T.ServiceError{
			Error:   docErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching user docs",
		}
	}

	marshalContent, marshalErr := json.Marshal(docBody.FileContent)
	if marshalErr != nil {
		return 0, &T.ServiceError{
			Error:   marshalErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while marshalling content",
		}
	}

	logs := M.Log{
		DocumentID:  doc.ID,
		PrevContent: doc.Content,
		NewContent:  null.JSONFrom(marshalContent),
		EditedBy:    docBody.EditedBy,
	}

	logsErr := logs.Insert(dbCtx, trx, boil.Infer())
	if logsErr != nil {
		return 0, &T.ServiceError{
			Error:   logsErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while inserting logs",
		}
	}

	doc.Content = null.JSONFrom(marshalContent)
	_, updateErr := doc.Update(dbCtx, trx, boil.Infer())

	if updateErr != nil {
		return 0, &T.ServiceError{
			Error:   updateErr,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while updating user docs",
		}
	}

	return 0, nil
}
