package model

import (
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseModel struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       *string            `json:"title" bson:"title" validate:"required,max=50"`
	Amount      *float64           `json:"amount" bson:"amount" validate:"required,number,gt=0"`
	Date        *string            `json:"date" bson:"date" validate:"datetime=02-01-2006"`
	Category    *string            `json:"category" bson:"category" validate:"required"`
	Description *string            `json:"description" bson:"description" validate:"required,max=20"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

func (em *ExpenseModel) TrimSpace() {
	val := reflect.ValueOf(em).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.String {
			field.Elem().SetString(strings.TrimSpace(field.Elem().String()))
		}
	}
}

func (im *ExpenseModel) MarshalBSON() ([]byte, error) {
	if im.CreatedAt.IsZero() {
		im.CreatedAt = time.Now()
	}
	im.UpdatedAt = time.Now()
	type my IncomeModel
	return bson.Marshal((*my)(im))
}
