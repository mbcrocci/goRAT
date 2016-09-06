package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1337")
	if err != nil {
		log.Fatalf("[ERROR] Can't create connection!: %v", err)
	}

	//Escrever na socket
	fmt.Fprintf(conn, "Testing connection")

	response, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("[RESP] - ", response)
}