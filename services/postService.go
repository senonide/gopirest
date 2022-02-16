/*
	This class is in charge of implementing the CRUD methods
*/
package services

import (
	"gopirest/gopirest/database"
	"gopirest/gopirest/models"
)

type PostService struct{}
type IPostService interface {
	Create(post models.Post)
	Read()
	ReadOne(postId string)
	Update(post models.Post, postId string)
	Delete(postId string)
}

func NewPostService() *PostService {
	var e *PostService
	return e
}

func (s *PostService) Create(post models.Post) error {

	err := database.Create(post)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) Read() ([]models.Post, error) {

	posts, err := database.Read()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) ReadOne(postId string) (models.Post, error) {

	post, err := database.ReadOne(postId)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (s *PostService) Update(post models.Post, postId string) error {

	err := database.Update(post, postId)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) Delete(postId string) error {

	err := database.Delete(postId)

	if err != nil {
		return err
	}

	return nil
}
