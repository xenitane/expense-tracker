package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": fmt.Sprintf("this path: `%v` does not exist."),
		"error":   errors.New("Address Not Found Error"),
	})
}
