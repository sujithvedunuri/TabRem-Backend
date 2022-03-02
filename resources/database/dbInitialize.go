package database

import (
	"fmt"
	"log"
	"os"
	"sujith/tabRemBackend/beans"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//"root:password@tcp(127.0.0.1:3306)/TabRemDatabase?charset=utf8mb4&parseTime=True&loc=Local"

	dbformat := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, error := gorm.Open(mysql.Open(dbformat), &gorm.Config{})
	Db = db
	if error != nil {
		log.Fatal("unable to connect to the database")
	}
	Db.AutoMigrate(&beans.Medicine{})

}
