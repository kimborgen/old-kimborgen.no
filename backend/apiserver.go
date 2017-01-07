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
	var activateTls bool
	prodPtr := flag.Bool("prod", false, "activate production enviroment")
	tlsPtr := flag.Bool("tls", false, "activate tls encryption")
	flag.Parse()
	if *prodPtr {
		log.Println("Production enviroment detected")
		production = true
	} else {
		log.Println("Development enviroment detected")
	}
	if *tlsPtr {
		log.Println("tls activated")
		activateTls = true
	} else {
		activateTls = false
	}

	dbStart()
	cleanUp()
	router := NewRouter()
	if production {
		host = "146.185.153.19:8090"
	} else {
		host = "localhost:8080"
	}
	if activateTls {
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

		srv := &http.Server{
			Addr:         host,
			Handler:      router,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Println("LetsGo!")
		log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	} else {
		log.Println("LetsGo!")
		log.Fatal(http.ListenAndServe(host, router))
	}
}

func cleanUp() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		dbClose()
		log.Println("cleaned up! goodbye")
		os.Exit(1)
	}()

}
