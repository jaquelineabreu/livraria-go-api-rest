package database

import (
	"time"
	"log"
	"gorm.io/gorm"
)

var db *gorm.DB


func StartDB(){
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil{
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)


}

func GetDatabase() *gorm.DB{
	return db
}




