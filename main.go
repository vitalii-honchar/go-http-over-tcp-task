package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const port = 8080

func main() {
	log.Printf("Starting TCP server at port %v\n", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalln("Can't start server", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "I see you connected\n")
}
