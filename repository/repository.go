package repository

import (
  "os"
  "fmt"

	"github.com/shimodii/dekamond-otp-challenge/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseConnection() {

  // docker :D
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

  // used from gorm docs :)
  //dsn := "host=localhost user=postgres password=password dbname=dekamond port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  
  if err != nil {
    panic("err")
  }

  db.AutoMigrate(&model.User{})
  

  Database = db
}
