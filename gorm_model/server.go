package gorm_model

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Run() {
	database := "yoyo_chat"
	tableName := "union_payouts_new"
	filename := tableName + ".go"
	modelPath := "./model"
	sql := "select COLUMN_NAME as name,DATA_TYPE as data_type,COLUMN_COMMENT as comment from information_schema.COLUMNS where TABLE_SCHEMA=? and  TABLE_NAME=?"
	columns := make([]Column, 0)
	db.Raw(sql, database, tableName).Scan(&columns)

	str := ""
	importStr := ""
	for _, v := range columns {
		if v.DataType == "date" || v.DataType == "datetime" || v.DataType == "timestamp" {
			importStr = `import "time"`
		}
		tag := fmt.Sprintf(`json:"%s" gorm:"column:%s"`, v.Name, v.Name)
		if _, ok := dataTypes[v.DataType]; ok {
			str += "\t" + translateName(v.Name) + "\t" + dataTypes[v.DataType] + "\t`" + tag + "`" + "\t//" + v.Comment + "\n"
		}
	}
	content := fmt.Sprintf(text, importStr, translateName(tableName), strings.TrimRight(str, "\n"), translateName(tableName), tableName)
	saveFile(modelPath, filename, content)
}

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"huoyu789",
		"47.104.186.50",
		13306,
		"yoyo_chat",
	)

	db, err = gorm.Open("mysql", connStr)
	if err != nil || db == nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3)
	db.DB().SetMaxOpenConns(20)
	//db = db.LogMode(true)
	db.DB().SetConnMaxLifetime(time.Duration(60) * time.Second)
}
