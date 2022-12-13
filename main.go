package main

import (
	"go-http-over-tcp-task/tcp"
	"log"
)

const port = 8080

func main() {
	s := tcp.NewTcpServer(port)
	err := s.Start()
	if err != nil {
		log.Fatalln("Can't start TCP server!", err)
	}
}
