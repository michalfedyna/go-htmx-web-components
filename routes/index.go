package routes

import "github.com/gofiber/fiber/v2"

func GetIndex(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		route := "routes/index"
		bind := fiber.Map{"PageTitle": "Welcome"}
		layout := "layouts/base"

		return c.Render(route, bind, layout)
	})
}
