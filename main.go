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
	"int":     " int",
	"varchar": " string",
}

func main() {
	tableName := "test"
	filename := tableName + ".go"
	modelPath := "./model"
	db := initdb()
	sql := "select COLUMN_NAME as name,DATA_TYPE as data_type from information_schema.COLUMNS where TABLE_NAME=?"
	columns := make([]Column, 0)
	db.Raw(sql, tableName).Scan(&columns)
	text := `package model

type %s struct{
%s
}

func (*%s) TableName() string {
	return "%s"
}
`
	str := ""
	for _, v := range columns {
		tag := fmt.Sprintf(`json:"%s" db:"%s"`, v.Name, v.Name)
		str += "\t" + translateName(v.Name) + dataTypes[v.DataType] + "\t`" + tag + "`" + "\n"
	}
	content := fmt.Sprintf(text, FirstCharToUpper(tableName), strings.TrimRight(str, "\n"), FirstCharToUpper(tableName), tableName)
	_, err := os.Stat(modelPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(modelPath, os.ModeDir)
		}
	}
	os.WriteFile(modelPath+"/"+filename, []byte(content), os.ModeAppend)
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
		"123456",
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
	db = db.LogMode(true)
	db.DB().SetConnMaxLifetime(time.Duration(60) * time.Second)
	return db
}
