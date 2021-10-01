package models

// To export the struct out of module "models", the first letter must be capitalzed
type Post struct {
	ID      string `bson:"_id,omitempty" json:"ID,omitempty"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type AllPosts []Post
