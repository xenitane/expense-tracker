package variables

import "go.mongodb.org/mongo-driver/mongo"

var (
	DB       *mongo.Client
	Database *mongo.Database
)
