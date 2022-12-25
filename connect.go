package database

import (
	"log"
	"inter/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connectme() {

	var err error

	db, err = gorm.Open("mysql", "root:Balaram@123@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Unable to connect to database")
	} else {
		log.Println("Connection Successful")
	}

	db.AutoMigrate(&model.Idlibatch{})

}
