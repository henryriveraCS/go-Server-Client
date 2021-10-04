package main

import (
	"fmt"
	"net"
	"log"
)

func handleConnection(conn net.Conn) {
	//tmp bytes to read from the connection
	tmp := make([]byte, 1024)
	for {
		r, err := conn.Read(tmp)
		if err != nil {
			break
		}
		log.Printf("Bytes received: %d", r)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		//handle the error here
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//handle error here
			fmt.Println("Error starting up: ", err)
		}
		go handleConnection(conn)
	}
}
