package routes

import "github.com/gofiber/fiber/v2"

type Page struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	ImageURL    string `json:"imgUrl"`
}

type HTTPResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Home",
		"Subtitle": "Just a collection of websites",
	}, "layouts/main")
}

func IndexJSONHandler(c *fiber.Ctx) error {
	return c.JSON(HTTPResponse{Message: "Yeah", Success: true,
		Data: []Page{
			{
				"https://isak-tech.tk",
				"Personal site for promoting my services",
				"/public/images/isaktech.png",
			},
			{
				"https://paste.isak-tech.tk",
				"Paste whatever you'd like and grab it later",
				"/public/images/paste.png",
			},
			{
				"https://mail.isak-tech.tk",
				"my personal mail server",
				"/public/images/email.png",
			},
			{
				"https://www.shapedivider.app/",
				"Create cool and awesome backgrounds",
				"/public/images/bgcreator.png",
			},
			{
				"https://www.cssmatic.com/box-shadow",
				"Experiment with the box shadow property",
				"/public/images/boxshadow.png",
			},
			{
				"https://tailwindcss.com/docs",
				"New interesting UI library",
				"/public/images/tw.png",
			},
			{
				"https://michalsnik.github.io/aos/",
				"An awesome animation library",
				"/public/images/aos.png",
			},
			{
				"https://www.iloveimg.com/",
				"A solid image compressor",
				"/public/images/imgcompress.png",
			},
			{
				"https://autoprefixer.github.io/",
				"Vendor Prefixer CSS/SCSS",
				"/public/images/vendorprefix.png",
			},
		},
	})
}
