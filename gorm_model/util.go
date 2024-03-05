package gorm_model

import (
	"os"
	"strings"
)

func saveFile(modelPath, filename, content string) {
	_, err := os.Stat(modelPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(modelPath, os.ModeDir)
		}
	}
	filePath := modelPath + "/" + filename
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(filePath, []byte(content), os.ModeAppend)
			return
		}
	}
	os.Remove(filePath)
	os.WriteFile(filePath, []byte(content), os.ModeAppend)
}

func translateName(name string) string {
	vs := strings.Split(name, "_")
	column := ""
	for _, val := range vs {
		column += FirstCharToUpper(val)
	}
	return column
}

func FirstCharToUpper(str string) string {
	return strings.ToUpper(string([]rune(str)[0])) + string([]rune(str)[1:])
}
