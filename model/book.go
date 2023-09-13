package model

import "fiber_postgres/database"

type Book struct {
	database.DefaultModel
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
