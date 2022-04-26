package handlers

import (
	"gopirest/gopirest/internal/models"
	"gopirest/gopirest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *services.PostService
}

type IPostHandler interface {
	GetPosts(*gin.Context, *services.PostService)
	GetPostById(*gin.Context, *services.PostService)
	CreatePost(*gin.Context, *services.PostService)
	UpdatePost(*gin.Context, *services.PostService)
	DeletePost(*gin.Context, *services.PostService)
}

func NewPostHandler() *PostHandler {
	eh := new(PostHandler)
	eh.service = services.NewPostService()
	return eh
}

func (eh *PostHandler) GetPosts(c *gin.Context) {

	Posts, err := eh.service.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, Posts)
}

func (eh *PostHandler) GetPostById(c *gin.Context) {
	id := c.Param("id")
	Post, err := eh.service.ReadOne(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, Post)
}

func (eh *PostHandler) CreatePost(c *gin.Context) {
	var newPost models.Post

	err := c.BindJSON(&newPost)
	if err != nil {
		return
	}

	eh.service.Create(newPost)

	c.IndentedJSON(http.StatusCreated, newPost)
}

func (eh *PostHandler) UpdatePost(c *gin.Context) {

	var newPost models.Post
	id := c.Param("id")

	err := c.BindJSON(&newPost)
	if err != nil {
		return
	}

	eh.service.Update(newPost, id)

	c.IndentedJSON(http.StatusOK, newPost)

}

func (eh *PostHandler) DeletePost(c *gin.Context) {

	id := c.Param("id")
	eh.service.Delete(id)

	Posts, err := eh.service.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, Posts)

}
