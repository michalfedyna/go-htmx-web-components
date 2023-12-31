package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"go-web/routes"
)

func bootstrap() *fiber.App {
	engine := html.New("./templates", ".html")
	app := fiber.New(
		fiber.Config{Views: engine})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(compress.New())
	app.Use(favicon.New(favicon.Config{
		File: "./static/img/favicon.ico",
		URL:  "static/img/favicon.ico",
	}))

	return app
}

func main() {
	app := bootstrap()

	app.Static("/static", "./static")
	routes.GetIndex(app)
	routes.GetSignIn(app)
	routes.GetGetStarted(app)
	routes.GetNotFound(app)

	app.Post("/click", func(c *fiber.Ctx) error {
		route := "partials/click"

		return c.Render(route, nil)
	})

	err := app.Listen(":3000")
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return
	}
}

// TODO: Modify CSS to shorter classes names
