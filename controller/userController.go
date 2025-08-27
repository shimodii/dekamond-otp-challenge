package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/dekamond-otp-challenge/model"
	"github.com/shimodii/dekamond-otp-challenge/repository"
	"github.com/shimodii/dekamond-otp-challenge/service"
)

func GetAllUsers(c *fiber.Ctx) error {
  users := []model.User{}

  limit := 10
  page := c.QueryInt("page", 1)
  
  page = max(1, page)

  offset := (page-1) * limit

  repository.Database.Limit(limit).Offset(offset).Find(&users)
  //fmt.Println(users)
  //repository.Database.Find(&users)

  return c.JSON(users)
}

func GetSingleUser(c *fiber.Ctx) error {
  id := c.Params("id")

  var user model.User

  repository.Database.Find(&user, id)

  return c.JSON(user)
}

func Phonenumber(c *fiber.Ctx) error {
  var user model.MinUser
  err := c.BodyParser(&user)
  if err != nil {
    panic(err)
  }

  code := service.GenCode(3)
  fmt.Println("new code genrated:", user.Phone, code)
  repository.RedisDB.Set(context.Background(), user.Phone, code, 2*time.Minute)

  return c.SendString("ok")
}

func OtpPhonenum(c *fiber.Ctx) error {
  var user model.SignUser
  err := c.BodyParser(&user)
  if err != nil {
    panic(err)
  }

  if service.ValidateCode(user.Code, user.Phone) {
    //fmt.Println("hellllloooo")
    var newUser model.User
    res := repository.Database.Where("phone = ?", user.Phone).First(&newUser)
    if res.Error != nil {
      var signUpUser model.User
      
      signUpUser.Phone = user.Phone
      signUpUser.CreatedAt = time.Now()
      repository.Database.Create(&signUpUser)
      return c.JSON(signUpUser)
    }
    return c.SendString("jwt token yay")
  }
  
  return c.SendString("no no no")
}
