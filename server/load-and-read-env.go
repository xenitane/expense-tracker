package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dotenv-org/godotenvvault"
	"github.com/xenitane/expense-tracker/server/variables"
)

func loadAndReadEnv() {
	var err error
	err = godotenvvault.Load()
	if err != nil {
		log.Fatal(err)
	}
	variables.PORT, err = strconv.Atoi(os.Getenv(PORT))
	if err != nil {
		log.Fatal(err)
	}
	if variables.PORT <= 0 {
		log.Fatal("inavlid port")
	}
	variables.MONGOURI = os.Getenv(MONGOURI)
	if variables.MONGOURI == "" {
		log.Fatal("empty mongouri")
	}
	variables.DBNAME = os.Getenv(DBNAME)
	if variables.DBNAME == "" {
		log.Fatal("empty dbname")
	}
}
