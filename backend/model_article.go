package main

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string
	Body  string
}
