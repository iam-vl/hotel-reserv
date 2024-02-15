package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/hotel-reserv/api"
)

func main() {
	// Pointer to a string
	listenAddr := flag.String("listenAddr", ":1111", "Listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	// app.Get("/foo", handleFoo)
	apiV1.Get("/user", api.HandleListUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)
	// Dereference the pointer
	app.Listen(*listenAddr)

}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Working just fine"})
}
func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Foo"})
}
