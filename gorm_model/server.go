package gorm_model

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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
