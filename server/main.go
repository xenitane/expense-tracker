package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/xenitane/expense-tracker/server/configs"
	"github.com/xenitane/expense-tracker/server/variables"
)

func main() {

	loadAndReadEnv() // loading the environment

	variables.App = fiber.New() // instantiating the web-server

	defer configs.ConnectDB()  // database-connection
	configs.MiddlewareConfig() // middleware-setup
	configs.RouteConfig()      // routing-setup

	log.Printf("Starting the web server at port: %v", variables.PORT) //starting the web-server

	if err := variables.App.Listen(fmt.Sprintf(":%v", variables.PORT)); err != nil {
		log.Fatalf("There was an error: %v\n while starting the server at port: %v", err, variables.PORT)
	}
}
