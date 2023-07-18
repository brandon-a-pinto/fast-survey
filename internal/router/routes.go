package router

import (
	"github.com/brandon-a-pinto/fast-survey/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func routes(a *fiber.App) {
	users := handler.UserHandler
	v1 := a.Group("/api/v1")
	{
		v1.Post("/users", users.HandlePostUser)
	}
}
