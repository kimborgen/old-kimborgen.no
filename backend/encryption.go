package main

import (
	"log"
	"net/http"

	"github.com/kabukky/httpscerts"
)

func GenerateTLSpair() {
	// Check if the cert files are available.
	// If they are not available, generate new ones.
	err := httpscerts.Check("cert.pem", "key.pem")
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "localhost:8080")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}
}

func TlsEncryption(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	})
}
