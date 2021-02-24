package routes

import "github.com/gofiber/fiber/v2"

func AboutHandler(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title":    "About",
		"Subtitle": "",
	}, "layouts/main")
}
