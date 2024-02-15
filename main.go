package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/foo", handleFoo)
	app.Listen(":1111")

}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Working just fine"})
}
