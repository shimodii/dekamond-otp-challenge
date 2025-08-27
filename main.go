package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/dekamond-otp-challenge/controller"
	"github.com/shimodii/dekamond-otp-challenge/model"
	"github.com/shimodii/dekamond-otp-challenge/repository"
)

func main() {
  application := fiber.New()

  ctx := context.Background()

  // database
  repository.DatabaseConnection()
  repository.OpenRedis()

  application.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello world")
  })

  // all tests are here
  /*
    i had no idea about redis so i opened this endpoint to test every single of my questions :D 
  */
  application.Get("/test", func(c *fiber.Ctx) error {
    user := model.User{
      Phone: "9933345821",
      CreatedAt: time.Now(),
    }
    repository.Database.Create(&user)
    repository.RedisDB.Set(ctx, user.Phone, "12345", 10*time.Second)
    return c.SendString("ok")
  })

  application.Get("/users", controller.GetAllUsers)
  application.Get("/users/:id", controller.GetSingleUser)
  application.Post("/get-code", controller.Phonenumber)

  fmt.Println("App is running on 1628")
  application.Listen(":1628")
}
