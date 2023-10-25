package routes

import "github.com/gofiber/fiber/v2"

func GetNotFound(app *fiber.App) {
	app.Get("*", func(c *fiber.Ctx) error {
		route := "routes/404"
		bind := fiber.Map{"PageTitle": "Page Not Found"}
		layout := "layouts/base"

		return c.Render(route, bind, layout)
	})
}
