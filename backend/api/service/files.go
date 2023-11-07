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
	U "github.com/word-extractor/word-extractor-apis/util"
)

type FileBody struct {
	UserID      int                    `json:"user_id"`
	FileName    string                 `json:"file_name" validate:"required"`
	FileContent map[string]interface{} `json:"file_content"`
	FileLink    string                 `json:"file_link"`
}
type FilesFilter struct {
	UserID int   `json:"user_id"`
	DocId  []int `json:"doc_id"`
}

func UploadFile(file *FileBody, trx *sql.Tx) (*M.Document, *T.ServiceError) {

	dbCtx := context.Background()

	file.FileContent = make(map[string]interface{})
	file.FileContent["pending"] = true

	// Marshal the content
	marshalContent, err := json.Marshal(file.FileContent)
	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while marshalling content",
		}
	}

	// Insert the file metadata into the database
	insertDoc := M.Document{
		Name:    file.FileName,
		Content: null.JSONFrom(marshalContent),
		Link:    null.StringFrom(file.FileLink),
		OwnerID: file.UserID,
	}

	inserterr := insertDoc.Insert(dbCtx, trx, boil.Infer())

	if inserterr != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while inserting document",
		}
	}
	go U.WordExtractor(file.FileLink, insertDoc.ID, file.UserID)

	return &insertDoc, nil

}

func GetFiles(fileFilter *FilesFilter, trx *sql.Tx) (*M.DocumentSlice, *T.ServiceError) {

	dbCtx := context.Background()

	var query []qm.QueryMod

	query = append(query, qm.Where("owner_id = ?", fileFilter.UserID))

	// docIds := make([]interface{}, len(fileFilter.DocId))
	// for i, id := range fileFilter.DocId {
	// 	docIds[i] = id
	// }
	// if len(fileFilter.DocId) > 0 {
	// 	query = append(query, qm.WhereIn("id in ?", docIds...))
	// }

	result, err := M.Documents(query...).All(dbCtx, trx)

	if err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Error while fetching documents",
		}
	}

	return &result, nil
}
