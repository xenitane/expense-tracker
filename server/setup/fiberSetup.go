package setup

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/config"
	"github.com/xenitane/expense-tracker/server/values"
)

func StartApplication() {
	values.App = fiber.New()

	config.MiddlewareConfig()
	config.RouteConfig()

	log.Printf("Starting the web server at port: %v", values.PORT)

	if err := values.App.Listen(fmt.Sprintf(":%v", values.PORT)); err != nil {
		log.Fatalf("There was an error: %v\n while starting the server at port: %v", err, values.PORT)
	}
}
