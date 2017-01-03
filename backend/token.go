package main

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

var my_secret_key []byte = []byte("yolo")

func getToken() string {

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_token, err := token.SignedString(my_secret_key)

	if err != nil {
		log.Fatal(err)
	}
	return signed_token
}

func checkToken(token_string string) bool {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return my_secret_key, nil
	})
	if err != nil {
		log.Fatal(err)
		return false
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		return true
	} else {
		log.Fatal(err)
		return false
	}
}
