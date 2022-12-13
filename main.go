package main

import (
	"fmt"
	"go-http-over-tcp-task/tcp"
	"log"
	"strings"
)

const port = 8080

func main() {
	s := tcp.NewTcpServer(port)
	sh := &SimpleHandler{}
	s.Add(sh.Handle)

	err := s.Start()
	if err != nil {
		log.Fatalln("Can't start TCP server!", err)
	}
}

type SimpleHandler struct {
	request string
	method  string
	path    string
	version string
}

func (sh *SimpleHandler) Handle(in string) (string, bool) {
	log.Printf("Input data: %v", in)

	if sh.method == "" {
		parts := strings.Fields(in)

		sh.method = parts[0]
		sh.path = parts[1]
		sh.version = parts[2]
	}

	sh.request += "\n" + in
	// log.Printf("%+v\n", sh)

	if in == "" {
		body := "I see you connected\n"
		resp := fmt.Sprintf(`HTTP/1.1 200 OK
Content-Length: %d
Content-Type: text/plain

%s`, len(body), body)
		log.Println("Send response")
		return resp, false
	}

	return "", true
}
