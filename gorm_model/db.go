package gorm_model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func init() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"huoyu789",
		"",
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
