package models

type Post struct {
	Id      string `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
