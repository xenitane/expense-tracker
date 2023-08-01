package main

import (
	"github.com/xenitane/expense-tracker/server/setup"
)

func main() {

	setup.LoadAndReadEnv() // loading the environment

	defer (setup.ConnectDB())() // database-connection

	setup.StartApplication() // instantiating and setting up the web-server
}
