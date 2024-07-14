package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	jsonData := map[string]string{"message": "Hello, World 👋!"}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Send(jsonValue)
    })

    app.Listen(":3000")
}