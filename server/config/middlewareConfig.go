package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MiddlewareConfig(FiberApp *fiber.App) {
	FiberApp.Use(cors.New(cors.Config{
		MaxAge:       300,
		AllowOrigins: strings.Join([]string{"https://*", "https://*"}, ","),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
			fiber.MethodHead,
		}, ","),
		AllowHeaders:     "*",
		AllowCredentials: false,
	})) // cors
	FiberApp.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed})) // compress
	FiberApp.Use(helmet.New())                                                  // helmet
	FiberApp.Use(logger.New())                                                  // logger
	FiberApp.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})
}
