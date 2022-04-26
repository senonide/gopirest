/*
	This class is in charge of interacting with the database, it applies the design pattern called fachade
*/
package database

import (
	"context"

	"gopirest/gopirest/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = GetCollection("posts")
var ctx = context.Background()

func Create(post models.Post) error {

	var err error

	_, err = collection.InsertOne(ctx, post)

	if err != nil {
		return err
	}

	return nil
}

func Read() ([]models.Post, error) {

	var posts []models.Post

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var post models.Post
		err = cur.Decode(&post)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func ReadOne(postId string) (models.Post, error) {

	var post models.Post

	oid, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.M{"_id": oid}

	cur := collection.FindOne(ctx, filter)

	err := cur.Decode(&post)

	if err != nil {
		return post, err
	}

	return post, nil
}

func Update(post models.Post, postId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(postId string) error {

	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(postId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
