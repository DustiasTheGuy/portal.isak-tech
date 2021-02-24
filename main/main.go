package main

import (
	"log"
	"portal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(cors.New())
	app.Static("/public", "./public")

	app.Get("/", routes.IndexHandler)
	app.Get("/about", routes.AboutHandler)
	app.Get("/data", routes.IndexJSONHandler)

	log.Fatal(app.Listen(":8083"))
}
