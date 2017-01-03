package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username       string `json:"username"`
	Name           string
	Email          string
	Role           string
	HashedPassword string `json:"hashed_password"`
}

type UserSess struct {
	gorm.Model

	User User
}

func (u User) String() string {
	return u.Name
}
