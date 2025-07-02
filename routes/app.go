package routes

import (
	"github.com/ambrizals/go-ddd-fiber/config"
	"github.com/ambrizals/go-ddd-fiber/modules/hello-world"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	api := humafiber.New(app, config.DefaultConfig("Dummy Project", "1.0.0"))

	hello_world.MakeHelloWorldHandler(&api)

	return app
}
