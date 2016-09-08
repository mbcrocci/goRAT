package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Create Http connection with TLS
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Println("[ERROR] - ", err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	tr := &http.Transport{
		TLSClientConfig:    &config,
		DisableCompression: true,
	}

	tlsClient := &http.Client{Transport: tr}

	resp, err := tlsClient.Get("https://127.0.0.1:1337")
	if err != nil {
		log.Println("[ERROR] - ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] - ", err)
	}

	fmt.Println("[BODY] - ", string(body))

}
