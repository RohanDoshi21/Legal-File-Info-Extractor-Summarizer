package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/boil"
	DB "github.com/word-extractor/word-extractor-apis/db"
	M "github.com/word-extractor/word-extractor-apis/my_models"
)

var pythonMutex = &sync.Mutex{}

func WordExtractor(path string, docId int, ownerID int) {
	dbCtx := context.Background()
	fileContent := make(map[string]interface{})
	fileContent["pending"] = false

	metaData := calculateMetaData(path)

	updateBody := M.Document{
		ID:      docId,
		Content: null.JSONFrom(metaData),
		OwnerID: ownerID,
		Link:    null.StringFrom(path),
		Name:    path,
	}

	_, updateErr := updateBody.Update(dbCtx, DB.PostgresConn, boil.Infer())
	if updateErr != nil {
		panic(updateErr)
	}
}

func calculateMetaData(filePath string) []byte {

	pythonMutex.Lock()
	defer pythonMutex.Unlock()

	fmt.Println("Python script started for file: ", filePath)

	dirPath, _ := os.Getwd()
	pythonScript := dirPath + "/util/model.py"
	cmd := exec.Command("python3", pythonScript, filePath)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return output
}

func UpdateFileSummary(doc *M.Document, kvPairs map[string]string) {
	pythonMutex.Lock()
	defer pythonMutex.Unlock()

	var filePath string

	for k, v := range kvPairs {
		filePath += "Q: " + k + " , A: " + v + " \n"
	}

	dirPath, _ := os.Getwd()
	pythonScript := dirPath + "/util/updateKVModel.py"
	cmd := exec.Command("python3", pythonScript, filePath)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

	kvPairs["summary"] = string(output)
	marshalContent, err := json.Marshal(kvPairs)
	if err != nil {
		log.Fatal(err)
	}

	updateBody := M.Document{
		ID:        doc.ID,
		Name:      doc.Name,
		UpdatedAt: null.TimeFrom(time.Now()),
		OwnerID:   doc.OwnerID,
		Content:   null.JSONFrom(marshalContent),
		Link:      doc.Link,
	}

	_, err = updateBody.Update(context.Background(), DB.PostgresConn, boil.Infer())
	if err != nil {
		panic(err)
	}
}
