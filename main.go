package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/dekamond-otp-challenge/model"
	"github.com/shimodii/dekamond-otp-challenge/repository"
)

func main() {
  application := fiber.New()

  // database
  repository.DatabaseConnection()

  application.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello world")
  })

  application.Get("/test", func(c *fiber.Ctx) error {
    user := model.User{
      Phone: "9933345821",
      CreatedAt: time.Now(),
    }
    repository.Database.Create(&user)
    return c.SendString("ok")
  })

  application.Listen(":1628")
}
