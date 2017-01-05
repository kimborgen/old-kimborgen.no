package main

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	User User
	jwt.StandardClaims
}

var my_secret_key []byte = []byte("yolo")

func getToken(user User) string {

	//dont send the hashed password in the token
	hashed_pass := user.HashedPassword
	user.HashedPassword = ""
	// Create the Claims
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "me",
		},
	}
	user.HashedPassword = hashed_pass

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_token, err := token.SignedString(my_secret_key)

	if err != nil {
		log.Fatal(err)
	}
	return signed_token
}

func checkToken(token_string string, customClaims *MyCustomClaims) int {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return my_secret_key, nil
	})

	if err != nil {
		log.Fatal(err)
	}
	if !token.Valid {
		return 1
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Printf("%v %v", claims.User.Username, claims.StandardClaims.ExpiresAt)
		*customClaims = *claims
	} else {
		log.Fatal(err)
	}
	return 0
}
