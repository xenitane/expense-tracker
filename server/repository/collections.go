package repository

import (
	"github.com/xenitane/expense-tracker/server/values"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection(collectionName string) *mongo.Collection {
	collection := values.Database.Collection(collectionName)
	return collection
}
