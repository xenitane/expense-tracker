package controllers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "app is up and running",
		"error":   nil,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": fmt.Sprintf("this path: `%v` does not exist.", c.Context().URI()),
		"error":   errors.New("address not found error").Error(),
	})
}
