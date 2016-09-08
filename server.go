package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader

	tlsServer := &http.Server{TLSConfig: config}
	tlsServer.Addr = ":1337"

	http.HandleFunc("/", handleConn)

	err = tlsServer.ListenAndServeTLS("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Println("[ERROR] - ", err)
	}
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GOT YOUR REQUEST")
}
