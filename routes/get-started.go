package routes

import "github.com/gofiber/fiber/v2"

func GetGetStarted(app *fiber.App) {
	app.Get("/get-started", func(c *fiber.Ctx) error {
		route := "routes/get-started"
		bind := fiber.Map{"PageTitle": "Get Started"}
		layout := "layouts/split-auth"

		return c.Render(route, bind, layout)
	})
}
