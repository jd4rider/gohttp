package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//func Main() {
//	app := fiber.New()
//
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString("Hello, World!")
//	})
//
//	app.Listen(":3000")
//}

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", HelloWorld)
}

func main() {
	app := fiber.New()
	SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
