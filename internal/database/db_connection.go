/*
	This class is in charge of connecting the application to the database in mongodb
	The password is loaded from a file that is not in the repository for security reasons.
	You can modify the connection variables to connect to another database
*/
package database

import (
	"context"
	"fmt"
	"gopirest/gopirest/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {

	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.DbUsr, config.DbPw, config.DbHost)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	return client.Database(config.DbName).Collection(collection)
}
