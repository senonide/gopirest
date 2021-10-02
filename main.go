package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/se-nonide/go-API-REST/models"
	"github.com/se-nonide/go-API-REST/services"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome to my GO API! Here you can make CRUD operations with posts!")
}

func createPost(w http.ResponseWriter, r *http.Request) {

	var newPost models.Post

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Post Data")
	}

	json.Unmarshal(reqBody, &newPost)

	services.Create(newPost)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func getPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := services.Read()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getOnePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, err := services.ReadOne(postID)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	var updatedPost models.Post

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedPost)

	services.Update(updatedPost, postID)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(updatedPost)
	fmt.Fprintf(w, "The post with ID %v has been updated successfully", postID)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	services.Delete(postID)

	fmt.Fprintf(w, "The post with ID %v has been deleted successfully", postID)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", indexRoute)
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", getOnePost).Methods("GET")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
