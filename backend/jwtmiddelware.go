package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Jwtmiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		this_route := mux.CurrentRoute(r)
		for _, route := range routes {
			if route.Name == this_route.GetName() && route.Clearence != 0 {
				body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
				if err != nil {
					log.Fatal(err)
				}
				if err_2 := r.Body.Close(); err_2 != nil {
					log.Fatal(err_2)
				}
				type DataStructure struct {
					Token string      `json:"token"`
					Data  interface{} `json:"data"`
				}
				var dataStructure DataStructure
				if err_3 := json.Unmarshal(body, &dataStructure); err_3 != nil {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(422) // unprocessable entity
					if errr := json.NewEncoder(w).Encode(err); errr != nil {
						log.Fatal(err_3)
					}
					return
				}
				//check if user can access
				var claims MyCustomClaims
				checkToken(dataStructure.Token, &claims)
				if claims.User.Clearance >= route.Clearence {
					newBody, err_4 := json.Marshal(dataStructure.Data)
					if err_4 != nil {
						log.Fatal(err_4)
					}
					r.Body = ioutil.NopCloser(strings.NewReader(string(newBody)))
					r.ContentLength = int64(len(newBody))
				} else {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(404)
					if errr := json.NewEncoder(w).Encode("Not enough security clearance sorry :("); errr != nil {
						log.Fatal(errr)
					}
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
