package main

import "github.com/jinzhu/gorm"

type Kollektiv struct {
	gorm.Model

	Name   string   `json:"name"`
	People []string `json:"people"`
	Rooms  []string `json:"rooms"`
	Order  []int64
}
