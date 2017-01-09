package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

func getJsonData(dataStructure interface{}, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &dataStructure); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errr := json.NewEncoder(w).Encode(err); errr != nil {
			log.Fatal(err)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	type LoginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var loginDetails LoginDetails
	getJsonData(&loginDetails, w, r)
	var user User
	db.Where("username = $1", loginDetails.Username).First(&user)
	if user.Username != "" {
		if production {
			err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginDetails.Password))
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println(user)
		pay := getToken(user)
		type Payload struct {
			Token string `json:"token"`
		}
		payload := Payload{Token: pay}
		log.Println(json.Marshal(payload))

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if erro := json.NewEncoder(w).Encode(payload); erro != nil {
			log.Fatal(erro)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(404) // unprocessable entity
		if err := json.NewEncoder(w).Encode("User not found"); err != nil {
			log.Fatal(err)
		}

	}
}

func Kollektivet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var kollektiv Kollektiv
	db.Where("name = $1", name).First(&kollektiv)
	if kollektiv.Name != "" {
		write_json(kollektiv, w)
	} else {
		write_error("kollektiv not found", w)
	}
}
func KollektivetNew(w http.ResponseWriter, r *http.Request) {
	var kollektiv Kollektiv
	getJsonData(&kollektiv, w, r)
	var kollektiv_2 []Kollektiv
	if kollektiv.Name != "" {
		db.Where("name = $1", kollektiv.Name).Find(&kollektiv_2)
		if len(kollektiv_2) > 0 {
			write_error("choose another name", w)
		} else {
			db.Create(&kollektiv)
		}
	} else {
		write_error("empty name", w)
	}
}

func KollektivetUpdate(w http.ResponseWriter, r *http.Request) {
	var kollektiv Kollektiv
	getJsonData(&kollektiv, w, r)
	if kollektiv.Name != "" {
		db.Model(&kollektiv).Updates(kollektiv)
	} else {
		write_error("empty name", w)
	}
}

func write_error(err string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(404) // unprocessable entity
	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Fatal(err)
	}
}

func write_success(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func write_json(thing interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK) // unprocessable entity
	if errr := json.NewEncoder(w).Encode(thing); errr != nil {
		log.Fatal(errr)
	}
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errr := json.NewEncoder(w).Encode(err); errr != nil {
			panic(err)
		}
	}

	t := "yolo"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
