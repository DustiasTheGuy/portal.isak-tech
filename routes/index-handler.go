package routes

import (
	"fmt"
	"portal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
	pages, err := models.GetAllPages()

	if err != nil {
		return c.JSON(HTTPResponse{
			Message: "Internal Server Error",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(HTTPResponse{Message: "Yeah", Success: true, Data: pages})
}

func NewPageHandler(c *fiber.Ctx) error {
	var p models.Page

	if err := c.BodyParser(&p); err != nil {
		fmt.Println(err)

		return c.JSON(HTTPResponse{
			Message: "Unable to parse body",
			Success: false,
			Data:    nil,
		})
	}

	if err := p.SaveNewPage(); err != nil {
		fmt.Println(err)

		return c.JSON(HTTPResponse{
			Message: "Unable to save post",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(HTTPResponse{
		Message: "Saved New Post",
		Success: true,
		Data:    nil,
	})
}

func RemovePageHandler(c *fiber.Ctx) error {
	pageID, err := strconv.ParseInt(c.Params("pageID"), 10, 64)

	if err != nil {
		return c.JSON(HTTPResponse{
			Message: "Invalid Parameter Recieved",
			Success: false,
			Data:    nil,
		})
	}

	if err := models.RemovePage(pageID); err != nil {
		return c.JSON(HTTPResponse{
			Message: "Page may have been moved or deletedf",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(HTTPResponse{
		Message: "Removed Post",
		Success: true,
		Data:    nil,
	})
}
