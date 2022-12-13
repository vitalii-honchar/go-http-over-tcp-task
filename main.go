package main

import (
	"go-http-over-tcp-task/tcp"
	"log"
)

const port = 8080

func main() {
	s := tcp.NewTcpServer(port)
	s.Add(SimpleHandler)

	err := s.Start()
	if err != nil {
		log.Fatalln("Can't start TCP server!", err)
	}
}

func SimpleHandler(in any) (any, bool) {
	log.Printf("Input data: %v", in)
	return []byte("I see you connected\n"), false
}
