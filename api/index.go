package handler

import (
	//"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
