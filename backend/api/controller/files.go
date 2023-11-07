package controller

import (
	"github.com/gofiber/fiber/v2"
	S "github.com/word-extractor/word-extractor-apis/api/service"
	H "github.com/word-extractor/word-extractor-apis/handler"
	U "github.com/word-extractor/word-extractor-apis/util"
)

func UploadFile(ctx *fiber.Ctx) error {

	// Get authenticated user details from body
	authUser := U.GetAuthUser(ctx)

	// Get the file body from the context
	// fileBody := ctx.Locals("body").(*S.FileBody)
	file, fileErr := ctx.FormFile("file")

	if fileErr != nil {
		return H.BuildError(ctx, "File is missing", fiber.StatusBadRequest, fileErr)
	}

	path, saveErr := U.SaveFile(ctx, file, authUser.Id)

	if saveErr != nil {
		return H.BuildError(ctx, "Failed to save file", fiber.StatusInternalServerError, saveErr)
	}

	fileBody := &S.FileBody{}
	// Get the transaction from the fiber context
	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	fileBody.FileName = file.Filename
	
	fileBody.FileLink = path
	fileBody.UserID = authUser.Id

	insertedFile, err := S.UploadFile(fileBody, pgTrx)

	if err != nil {
		return H.BuildError(ctx, err.Message, err.Code, err.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":   1,
		"file": insertedFile,
	})

}

func GetFiles(ctx *fiber.Ctx) error {

	// Get authenticated user details from body
	authUser := U.GetAuthUser(ctx)

	// Get the transaction from the fiber context
	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	filesFilter := &S.FilesFilter{}
	filesFilter.UserID = authUser.Id

	// // Get the files id from the query parameters
	// if files := ctx.Query(("files_id")); files != "" {
	// 	fileList := U.StringToIntList(files)
	// 	filesFilter.DocId = fileList
	// }

	fileData, err := S.GetFiles(filesFilter, pgTrx)

	if err != nil {
		return H.BuildError(ctx, err.Message, err.Code, err.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"files": fileData,
	})

}
