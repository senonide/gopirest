/*
	In this class are defined all the models with which the application works
*/
package models

// To export the struct out of module "models" in go, the first letter must be a capital letter
type Post struct {
	ID      string `bson:"_id,omitempty" json:"ID,omitempty"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type AllPosts []Post
