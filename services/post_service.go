/*
	This class is in charge of implementing the CRUD methods
*/
package services

import (
	"github.com/se-nonide/go-API-REST/database"
	"github.com/se-nonide/go-API-REST/models"
)

func Create(post models.Post) error {

	err := database.Create(post)

	if err != nil {
		return err
	}

	return nil
}

func Read() (models.AllPosts, error) {

	posts, err := database.Read()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func ReadOne(postID string) (models.Post, error) {

	post, err := database.ReadOne(postID)

	if err != nil {
		return post, err
	}

	return post, nil
}

func Update(post models.Post, postID string) error {

	err := database.Update(post, postID)

	if err != nil {
		return err
	}

	return nil
}

func Delete(postID string) error {

	err := database.Delete(postID)

	if err != nil {
		return err
	}

	return nil
}
