package repository

import (
	"errors"

	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/value"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExpenses() ([]model.ExpenseModel, error) {
	var expense model.ExpenseModel
	var expenses []model.ExpenseModel = make([]model.ExpenseModel, 0)

	expenseCollection := getCollection(value.ConstExpenseCollection)
	cursor, err := expenseCollection.Find(value.Ctx, bson.D{}, &options.FindOptions{
		Sort: bson.D{{Key: "created_at", Value: -1}},
	})
	if err != nil {
		defer cursor.Close(value.Ctx)
		return expenses, err
	}
	for cursor.Next(value.Ctx) {
		err := cursor.Decode(&expense)
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func GetExpenseByID(expenseID string) (model.ExpenseModel, error) {
	var expense model.ExpenseModel
	expenseCollection := getCollection(value.ConstExpenseCollection)
	objectID, err := primitive.ObjectIDFromHex(expenseID)
	if err != nil {
		return expense, err
	}
	err = expenseCollection.FindOne(value.Ctx, bson.D{{Key: "_id", Value: objectID}}).Decode(&expense)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func SaveExpense(expense *model.ExpenseModel) (model.ExpenseModel, error) {
	var expenseRes model.ExpenseModel
	expenseCollection := getCollection(value.ConstExpenseCollection)
	result, err := expenseCollection.InsertOne(value.Ctx, expense)
	if err != nil {
		return expenseRes, err
	}
	expenseRes, err = GetExpenseByID(result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return expenseRes, err
	}
	return expenseRes, nil
}

func DeleteExpense(id string) error {
	expenseCollection := getCollection(value.ConstExpenseCollection)
	result, err := expenseCollection.DeleteOne(value.Ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("no record with this id exists in the database")
	}
	return nil
}
