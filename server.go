package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("[ERROR] - Establishing server: %v", err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal("[ERROR] - Can't accept connections: %v", err)
	}

	//log.Println("[REQ] - got: ", request)

	conn.Write([]byte("Connection successful!"))
}
