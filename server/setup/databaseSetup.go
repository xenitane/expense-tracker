package setup

import (
	"log"

	"github.com/xenitane/expense-tracker/server/values"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() func() {
	var err error
	values.DB, err = mongo.Connect(values.Ctx, options.Client().ApplyURI(values.MONGOURI))
	if err != nil {
		log.Fatal(err)
	}
	disconnect := func() {
		if err := values.DB.Disconnect(values.Ctx); err != nil {
			panic(err)
		}
	}
	values.Database = values.DB.Database(values.DBNAME)

	if err = values.DB.Ping(values.Ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return disconnect
}
