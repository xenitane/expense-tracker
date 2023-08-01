package repository

import (
	"sort"

	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/values"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getIncomeCollection() *mongo.Collection {
	if _, isPresent := values.Collections[values.C_INCOMECOLLECTION]; !isPresent {
		values.Collections[values.C_INCOMECOLLECTION] = getCollection(values.C_INCOMECOLLECTION)
	}
	return values.Collections[values.C_INCOMECOLLECTION]
}

func GetIncomes() ([]model.IncomeModel, error) {
	var income model.IncomeModel
	var incomes []model.IncomeModel

	incomeCollection := getIncomeCollection()
	cursor, err := incomeCollection.Find(values.Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(values.Ctx)
		return incomes, err
	}
	for cursor.Next(values.Ctx) {
		err := cursor.Decode(&income)
		if err != nil {
			sort.SliceStable(incomes, func(i, j int) bool {
				return incomes[i].CreatedAt.Before(incomes[j].CreatedAt)
			})
			return incomes, err
		}
		incomes = append(incomes, income)
	}
	sort.SliceStable(incomes, func(i int, j int) bool {
		return incomes[i].CreatedAt.Before(incomes[j].CreatedAt)
	})
	return incomes, nil
}

func GetIncomeByID(incomeID string) (model.IncomeModel, error) {
	var income model.IncomeModel
	incomeCollection := getIncomeCollection()
	objectID, err := primitive.ObjectIDFromHex(incomeID)
	if err != nil {
		return income, err
	}
	err = incomeCollection.FindOne(values.Ctx, bson.D{{Key: "_id", Value: objectID}}).Decode(&income)
	if err != nil {
		return income, err
	}
	return income, nil
}

func SaveIncome(income *model.IncomeModel) (model.IncomeModel, error) {
	var incomeRes model.IncomeModel
	incomeCollection := getIncomeCollection()
	result, err := incomeCollection.InsertOne(values.Ctx, income)
	if err != nil {
		return incomeRes, err
	}
	incomeRes, err = GetIncomeByID(result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return incomeRes, err
	}
	return incomeRes, nil
}

func UpdateIncome() {}
