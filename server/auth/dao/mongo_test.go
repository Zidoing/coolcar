package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestMongo_ResolveAccountID(t *testing.T) {
	ctx := context.Background()
	mc, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.50.5:27017/?readPreference=primary&appname=mongodb-vscode%200.6.10&ssl=false"))
	if err != nil {
		panic(err)
	}
	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(ctx, "123")
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
