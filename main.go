package main

import (
	"task-management-system/configs"
	"task-management-system/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.EnvMongoURI()
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.BaseRoute(app)
	app.Listen(":6000")
}
