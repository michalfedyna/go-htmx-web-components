package routes

import "github.com/gofiber/fiber/v2"

func GetSignIn(app *fiber.App) {
	app.Get("/sign-in", func(c *fiber.Ctx) error {
		route := "routes/sign-in"
		bind := fiber.Map{"PageTitle": "Sign In"}
		layout := "layouts/split-auth"

		return c.Render(route, bind, layout)
	})
}
