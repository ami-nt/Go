package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    Book   `json:"data"`
	Message string `json:"message"`
	Datas   []Book `json:"datas"`
}
