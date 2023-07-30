package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/xenitane/expense-tracker/server/variables"
)

func ConnectDB() func() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	variables.DB, err = mongo.Connect(ctx, options.Client().ApplyURI(variables.MONGOURI))
	if err != nil {
		log.Fatal(err)
	}
	disconnect := func() {
		if err := variables.DB.Disconnect(ctx); err != nil {
			panic(err)
		}
	}

	variables.Database = variables.DB.Database(variables.DBNAME)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = variables.DB.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return disconnect
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := variables.Database.Collection(collectionName)
	return collection
}
