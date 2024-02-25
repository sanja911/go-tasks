package routes

import (
	userControllers "task-management-system/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func BaseRoute(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")
	user.Post("/", userControllers.Register)
	user.Post("/login", userControllers.Login)
}
