package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var production bool = false
var host string

func main() {
	boolPtr := flag.Bool("prod", false, "activate production enviroment")
	flag.Parse()
	if *boolPtr {
		log.Println("Production enviroment detected")
		production = true
	} else {
		log.Println("Development enviroment detected")
	}
	dbStart()
	cleanUp()
	GenerateTLSpair()
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	router := NewRouter()

	if production {
		host = "146.185.153.19:22"
	} else {
		host = "localhost:80"
	}
	srv := &http.Server{
		Addr:         host,
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Println("LetsGo!")
	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
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
