package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	
	"github.com/yamagn/qbox/model"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    db, err = gorm.Open("sqlite3", "qbox.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
	}
	
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

func autoMigration() {
	db.AutoMigrate(&model.Post{})
}