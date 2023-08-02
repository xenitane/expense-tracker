package setup

import (
	"log"
	"os"
	"strconv"

	"github.com/dotenv-org/godotenvvault"
	"github.com/xenitane/expense-tracker/server/value"
)

func LoadAndReadEnv() {
	var err error
	err = godotenvvault.Load()
	if err != nil {
		log.Fatal(err)
	}
	value.Env[value.ConstPort], err = strconv.Atoi(os.Getenv(value.ConstPort))
	if err != nil {
		log.Fatal(err)
	}
	if value.Env[value.ConstPort].(int) <= 0 {
		log.Fatal("inavlid port")
	}
	value.Env[value.ConstMongoURI] = os.Getenv(value.ConstMongoURI)
	if value.Env[value.ConstMongoURI].(string) == "" {
		log.Fatal("empty mongouri")
	}
	value.Env[value.ConstDBName] = os.Getenv(value.ConstDBName)
	if value.Env[value.ConstDBName].(string) == "" {
		log.Fatal("empty dbname")
	}
}
