package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type allBooks []Book

var Books = allBooks{
	{
		ID:          1,
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Author:      "Golang, Inc",
	},
}
