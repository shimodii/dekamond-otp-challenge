package main

import "github.com/gofiber/fiber/v2"

func main() {
  application := fiber.New()

  application.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello world")
  })

  application.Listen(":1628")
}
