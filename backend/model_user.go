package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username       string `json:"username"`
	Name           string
	Email          string
	HashedPassword string `json:"hashed_password"`
	Clearance      int16
}
