package util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/boil"
	DB "github.com/word-extractor/word-extractor-apis/db"
	M "github.com/word-extractor/word-extractor-apis/my_models"
)

func WordExtractor(path string, docId int, ownerID int) {
	dbCtx := context.Background()
	fileContent := make(map[string]interface{})
	fileContent["pending"] = false

	metaData := calculateMetaData(path)
	var metadata2 FileMetadata
	err := json.Unmarshal(metaData, &metadata2)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	updateBody := M.Document{
		ID:      docId,
		Content: null.JSONFrom(metaData),
		OwnerID: ownerID,
		Link:    null.StringFrom(path),
		Name:    metadata2.FileName,
	}

	_, updateErr := updateBody.Update(dbCtx, DB.PostgresConn, boil.Infer())
	if updateErr != nil {
		panic(updateErr)
	}
}

type FileMetadata struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	Summary  string `json:"summary"`
}

func calculateMetaData(filePath string) []byte {
	dirPath, _ := os.Getwd()
	pythonScript := dirPath + "/util/summary.py"
	cmd := exec.Command("python3", pythonScript, filePath)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	// var metadata FileMetadata
	// err = json.Unmarshal(output, &metadata)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return output
}
