package gorm_model

import "github.com/jinzhu/gorm"

type Column struct {
	Name     string `json:"name"`
	DataType string `json:"data_type"`
	Comment  string `json:"comment"`
}

var dataTypes = map[string]string{
	"tinyint":   "int64",
	"bigint":    "int64",
	"int":       "int64",
	"varchar":   "string",
	"char":      "string",
	"text":      "string",
	"datetime":  "time.Time",
	"date":      "time.Time",
	"decimal":   "float64",
	"timestamp": "time.Time",
}

var db *gorm.DB
var err error
