package main

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type BookReqRes struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
