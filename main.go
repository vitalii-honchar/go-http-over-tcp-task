package main

import (
	"fmt"
	"go-http-over-tcp-task/tcp"
	"log"
	"strings"
)

const port = 8080

const responseBodyTemplate = `<h1>You made a request</h1> 
<p>method = %s</p> 
<p>path = %s</p> 
<p>version = %s</p>`

const responseTemplate = `HTTP/1.1 200 OK
Content-Length: %d
Content-Type: text/html

%s
`

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
	if sh.method == "" {
		parts := strings.Fields(in)

		sh.method = parts[0]
		sh.path = parts[1]
		sh.version = parts[2]
		log.Printf("method = %s, path = %s, version = %s\n", sh.method, sh.path, sh.version)
	}

	sh.request += "\n" + in

	if in == "" {
		body := fmt.Sprintf(responseBodyTemplate, sh.method, sh.path, sh.version)
		resp := fmt.Sprintf(responseTemplate, len(body), body)
		log.Printf("%s\n", sh.request)
		return resp, false
	}

	return "", true
}
