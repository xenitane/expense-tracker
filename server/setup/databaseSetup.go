package setup

import (
	"log"

	"github.com/xenitane/expense-tracker/server/value"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() func() {
	var err error
	value.DBClient, err = mongo.Connect(value.Ctx, options.Client().ApplyURI(value.Env[value.ConstMongoURI].(string)))
	if err != nil {
		log.Fatal(err)
	}
	disconnect := func() {
		if err := value.DBClient.Disconnect(value.Ctx); err != nil {
			panic(err)
		}
	}
	value.Database = value.DBClient.Database(value.Env[value.ConstDBName].(string))

	if err = value.DBClient.Ping(value.Ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return disconnect
}
