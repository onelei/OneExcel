package fileUtil

import (
	"OneExcel/src"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func DeleteDir(dirPath string) {
	os.RemoveAll(dirPath)
}

func DeleteFile(fileName string) {
	var err = os.Remove(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateDir(dirPath string) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFile(filePath string) {
	if FileExist(filePath) {
		DeleteFile(filePath)
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func IsExcelClosed() bool {
	hasExcelOpen := false
	files, _ := ioutil.ReadDir(src.EXCEL_DIR)
	for _, f := range files {
		fileName := f.Name()
		if strings.HasPrefix(fileName, "~") {
			hasExcelOpen = true
			log.Fatalf("Please Close Excel: %s", fileName)
		}
	}
	if hasExcelOpen {
		return false
	}
	return true
}

func GetAllExcelName() []string {
	var results []string
	files, _ := ioutil.ReadDir(src.EXCEL_DIR)
	for _, f := range files {
		fileName := f.Name()
		if strings.HasSuffix(fileName, ".xlsx") || strings.HasSuffix(fileName, ".xls") {
			results = append(results, fileName)
		}
	}
	return results
}
