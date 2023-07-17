package router

import "github.com/gofiber/fiber/v2"

func Start() {
	app := fiber.New()

	routes(app)

	app.Listen(":3000")
}
