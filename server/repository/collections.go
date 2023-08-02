package repository

import (
	"github.com/xenitane/expense-tracker/server/value"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection(collectionName string) *mongo.Collection {
	if _, isPresent := value.Collections[collectionName]; !isPresent {
		value.Collections[collectionName] = value.Database.Collection(collectionName)
	}
	return value.Collections[collectionName]
}
