package repository

import (
	"errors"

	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/value"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetIncomes() ([]model.IncomeModel, error) {
	var incomes []model.IncomeModel = make([]model.IncomeModel, 0)

	incomeCollection := getCollection(value.ConstIncomeCollection)
	cursor, err := incomeCollection.Find(value.Ctx, bson.M{}, &options.FindOptions{
		Sort: bson.D{{Key: "created_at", Value: -1}},
	})
	if err != nil {
		defer cursor.Close(value.Ctx)
		return incomes, err
	}
	if err = cursor.All(value.Ctx, &incomes); err != nil {
		return incomes, err
	}
	return incomes, nil
}

func GetIncomeByID(incomeID string) (model.IncomeModel, error) {
	var income model.IncomeModel
	incomeCollection := getCollection(value.ConstIncomeCollection)
	objectID, err := primitive.ObjectIDFromHex(incomeID)
	if err != nil {
		return income, err
	}
	err = incomeCollection.FindOne(value.Ctx, bson.M{"_id": objectID}).Decode(&income)
	if err != nil {
		return income, err
	}
	return income, nil
}

func SaveIncome(income *model.IncomeModel) (model.IncomeModel, error) {
	var incomeRes model.IncomeModel
	incomeCollection := getCollection(value.ConstIncomeCollection)
	result, err := incomeCollection.InsertOne(value.Ctx, income)
	if err != nil {
		return incomeRes, err
	}
	incomeRes, err = GetIncomeByID(result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return incomeRes, err
	}
	return incomeRes, nil
}

func DeleteIncome(id primitive.ObjectID) error {
	incomeCollection := getCollection(value.ConstIncomeCollection)
	result, err := incomeCollection.DeleteOne(value.Ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("no record with this id exists in the database")
	}
	return nil
}
