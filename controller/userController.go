package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/dekamond-otp-challenge/model"
	"github.com/shimodii/dekamond-otp-challenge/repository"
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
