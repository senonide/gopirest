package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "sergio"
	pwd      = "BbJ7by8ERQunZ7n"
	host     = "cluster0.qpvi4.mongodb.net"
	database = "nodeAngular"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", usr, pwd, host)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
