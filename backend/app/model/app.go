package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"fmt"
	"time"
)

var db *gorm.DB

func DB() *gorm.DB {
	// Ping test
	err := db.DB().Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func InitDatabase() {
	var err error

	var db_name = "stock"
	var db_user = "stockuser"
	var db_pw = "stockuser123"
	var db_host = "127.0.0.1"
	var db_port = 3306

	sqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db_user, db_pw, db_host, db_port, db_name)

	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		revel.AppLog.Error("Fail To Connection DB")
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
}
