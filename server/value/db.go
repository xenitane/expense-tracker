package value

import (
	"context"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBClient    *mongo.Client                = nil
	Database    *mongo.Database              = nil
	Collections map[string]*mongo.Collection = make(map[string]*mongo.Collection)
	Ctx         context.Context              = context.TODO()
	Validator   validator.Validate           = *validator.New()
)
