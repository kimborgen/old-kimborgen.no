package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var production bool = false

func main() {
	dbStart()
	cleanUp()
	log.Println("db is ready for dedi")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func cleanUp() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		destroyTables()
		dbClose()
		log.Println("cleaned up! goodbye")
		os.Exit(1)
	}()

}
