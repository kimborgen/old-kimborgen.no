package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("build"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe("146.185.153.19:80", nil)
}
