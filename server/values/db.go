package values

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB          *mongo.Client                = nil
	Database    *mongo.Database              = nil
	Collections map[string]*mongo.Collection = make(map[string]*mongo.Collection)
	Ctx         context.Context              = context.TODO()
)
