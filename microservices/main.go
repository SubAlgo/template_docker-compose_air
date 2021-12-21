package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := "OK"
		return c.SendString(msg)
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		msg := "about"
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":8181"))
}
