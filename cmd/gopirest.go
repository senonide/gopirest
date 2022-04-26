package main

import (
	"gopirest/gopirest/internal/config"
	"gopirest/gopirest/internal/handlers"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	postHandler := handlers.NewPostHandler()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/posts", postHandler.GetPosts)
	router.GET("/posts/:id", postHandler.GetPostById)
	router.POST("/posts", postHandler.CreatePost)
	router.PUT("/posts/:id", postHandler.UpdatePost)
	router.DELETE("/posts/:id", postHandler.DeletePost)

	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	var url string = "localhost:" + strconv.Itoa(config.Port)

	log.Printf("Listening on port: %d", config.Port)
	router.Run(url)
}
