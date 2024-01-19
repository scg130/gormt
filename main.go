package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Column struct {
	Name     string `json:"name"`
	DataType string `json:"data_type"`
}

var dataTypes = map[string]string{
	"tinyint":  " int8",
	"bigint":   " int64",
	"int":      " int",
	"varchar":  " string",
	"char":     " string",
	"text":     " string",
	"datetime": " time.Time",
	"date":     " time.Time",
	"decimal":  "float64",
}

func main() {
	tableName := "union_apply_new"
	filename := tableName + ".go"
	modelPath := "./model"
	db := initdb()
	sql := "select COLUMN_NAME as name,DATA_TYPE as data_type from information_schema.COLUMNS where TABLE_NAME=?"
	columns := make([]Column, 0)
	db.Raw(sql, tableName).Scan(&columns)
	text := `package model

%s

type %s struct{
%s
}

func (m *%s) TableName() string {
	return "%s"
}
`
	str := ""
	importStr := ""
	for _, v := range columns {
		if v.DataType == "date" || v.DataType == "datetime" {
			importStr = `import "time"`
		}
		tag := fmt.Sprintf(`json:"%s" gorm:"column:%s"`, v.Name, v.Name)
		if _, ok := dataTypes[v.DataType]; ok {
			str += "\t" + translateName(v.Name) + dataTypes[v.DataType] + "\t`" + tag + "`" + "\n"
		}
	}
	content := fmt.Sprintf(text, importStr, translateName(tableName), strings.TrimRight(str, "\n"), translateName(tableName), tableName)
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

func initdb() *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"smd013012",
		"192.168.1.130",
		3306,
		"test",
	)

	db, err := gorm.Open("mysql", connStr)
	if err != nil || db == nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3)
	db.DB().SetMaxOpenConns(20)
	//db = db.LogMode(true)
	db.DB().SetConnMaxLifetime(time.Duration(60) * time.Second)
	return db
}
