package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	routes(app)

	app.Listen(os.Getenv("LISTEN_ADDRESS"))
}
