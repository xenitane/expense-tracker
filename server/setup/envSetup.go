package setup

import (
	"log"
	"os"
	"strconv"

	"github.com/dotenv-org/godotenvvault"
	"github.com/xenitane/expense-tracker/server/values"
)

func LoadAndReadEnv() {
	var err error
	err = godotenvvault.Load()
	if err != nil {
		log.Fatal(err)
	}
	values.PORT, err = strconv.Atoi(os.Getenv(values.C_PORT))
	if err != nil {
		log.Fatal(err)
	}
	if values.PORT <= 0 {
		log.Fatal("inavlid port")
	}
	values.MONGOURI = os.Getenv(values.C_MONGOURI)
	if values.MONGOURI == "" {
		log.Fatal("empty mongouri")
	}
	values.DBNAME = os.Getenv(values.C_DBNAME)
	if values.DBNAME == "" {
		log.Fatal("empty dbname")
	}
}
