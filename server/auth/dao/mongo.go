package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{col: db.Collection("account")}
}

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{"open_id": openID}, bson.M{
		"$set": bson.M{
			"open_id": openID,
		},
	}, options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))

	err := res.Err()
	if err != nil {
		return "", fmt.Errorf("cannot findoneandupdate %v", err)
	}
	var row struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err = res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode %v", err)
	}
	fmt.Println(row)

	return row.ID.Hex(), err
}
