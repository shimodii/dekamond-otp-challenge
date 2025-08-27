package repository

import (
	"github.com/shimodii/dekamond-otp-challenge/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseConnection() {

  // used from gorm docs :)
  dsn := "host=localhost user=postgres password=password dbname=dekamond port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  
  if err != nil {
    panic("err")
  }

  db.AutoMigrate(&model.User{})
  

  Database = db
}
