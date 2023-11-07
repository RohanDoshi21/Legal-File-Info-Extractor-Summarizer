package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	S "github.com/word-extractor/word-extractor-apis/api/service"
	H "github.com/word-extractor/word-extractor-apis/handler"
	U "github.com/word-extractor/word-extractor-apis/util"
)

func GetUsers(ctx *fiber.Ctx) error {
	userFilter := &S.UserFilter{}
	pgTrx := U.GetPGTrxFromFiberCtx(ctx)
	if users := ctx.Query("users"); users != "" {
		userIds := U.StringToIntList(users)
		userFilter.Id = userIds
	}
	userData, err := S.GetUsers(userFilter, pgTrx)

	if err != nil {
		return H.BuildError(ctx, err.Message, err.Code, err.Error)
	}
	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"users": userData,
	})
}

func GetUserDocs(ctx *fiber.Ctx) error {
	userId := ctx.Query("user_id")
	if userId == "" {
		return H.BuildError(ctx, "user_id is required", fiber.ErrBadRequest.Code, nil)
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return H.BuildError(ctx, "Error converting", fiber.ErrInternalServerError.Code, err)
	}
	userDocFilter := &S.UserDocFilter{}
	userDocFilter.UserID = userIdInt
	pgTrx := U.GetPGTrxFromFiberCtx(ctx)
	userDocs, userDocErr := S.GetUserDocs(userDocFilter, pgTrx)
	if userDocErr != nil {
		return H.BuildError(ctx, userDocErr.Message, userDocErr.Code, userDocErr.Error)
	}
	return H.Success(ctx, fiber.Map{
		"ok":   1,
		"docs": userDocs,
	})
}

func UpdateUserDoc(ctx *fiber.Ctx) error {
	docBody := &S.DocBody{}
	if err := ctx.BodyParser(docBody); err != nil {
		msg := "Failed to parse the body!"
		return H.BuildError(ctx, msg, fiber.ErrBadRequest.Code, err)
	}
	docId := ctx.Query("doc_id")
	if docId == "" {
		return H.BuildError(ctx, "doc_id is required", fiber.ErrBadRequest.Code, nil)
	}
	docIdInt, err := strconv.Atoi(docId)
	if err != nil {
		return H.BuildError(ctx, "Error converting", fiber.ErrInternalServerError.Code, err)
	}

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	docBody.DocID = docIdInt

	authUser := U.GetAuthUser(ctx)
	docBody.EditedBy = authUser.Id

	_, updateErr := S.UpdateUserDoc(docBody, pgTrx)
	if updateErr != nil {
		return H.BuildError(ctx, updateErr.Message, updateErr.Code, updateErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":     1,
		"succes": "Document updated successfully",
	})
}
