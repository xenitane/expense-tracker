package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncomeModel struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       *string            `json:"title" bson:"title" validate:"required,max=50"`
	Amount      *float64           `json:"amount" bson:"amount" validate:"required,number,gt=0"`
	IncomeType  *string            `json:"type,omitempty" bson:"type,omitempty"`
	Date        *string            `json:"date" bson:"date" validate:"datetime=02-01-2006"`
	Category    *string            `json:"category" bson:"category" validate:"required"`
	Description *string            `json:"description" bson:"description" validate:"required,max=20"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

func (im *IncomeModel) MarshalBSON() ([]byte, error) {
	if im.CreatedAt.IsZero() {
		im.CreatedAt = time.Now()
	}
	im.UpdatedAt = time.Now()
	type my IncomeModel
	return bson.Marshal((*my)(im))
}
