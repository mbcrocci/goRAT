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
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("[ERROR] - Can't accept connections: %v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("[Error] - Can't read: ", err)
	}
	log.Println("[REQUEST] - ", string(buf))
	conn.Write([]byte("Connection Successful"))
	conn.Close()
}
