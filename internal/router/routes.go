package router

import "github.com/gofiber/fiber/v2"

func routes(a *fiber.App) {
	v1 := a.Group("/api/v1")
	{
		v1.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"msg": "Hello World",
			})
		})
	}
}
