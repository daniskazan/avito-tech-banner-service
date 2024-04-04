package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
