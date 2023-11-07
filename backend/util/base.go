package util

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	T "github.com/word-extractor/word-extractor-apis/types"
	"golang.org/x/crypto/bcrypt"
)

func GetAuthUser(ctx *fiber.Ctx) *T.AuthUser {
	return ctx.Locals("user").(*T.AuthUser)
}

func GetPGTrxFromFiberCtx(ctx *fiber.Ctx) *sql.Tx {
	trxInf := ctx.Locals("pgTrx")

	if trxInf == nil {
		return nil
	}

	return ctx.Locals("pgTrx").(*sql.Tx)
}

// Return the hash of the password
func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare the password with the stored password hash
func ComparePassword(storedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}

func StringToIntList(str string) []int {
	filesList := strings.Split(str, ",")

	list := make([]int, 0)
	for _, file := range filesList {
		fileID, err := strconv.Atoi(file)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, fileID)
	}
	return list
}

func SaveFile(ctx *fiber.Ctx,file *multipart.FileHeader, userId int) (string,error){

	path := "./uploads/" + fmt.Sprint(userId) + "/"
	_, pathErr := os.Stat(path)
	if os.IsNotExist(pathErr) {
		os.MkdirAll(path, os.ModePerm)
	}

	path += file.Filename
	saveErr := ctx.SaveFile(file, path)

	return path ,saveErr
}
