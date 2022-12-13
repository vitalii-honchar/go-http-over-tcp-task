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

const indexBodyTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>GET INDEX</title>
</head>
<body>
	<h1>"GET INDEX"</h1>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
	<h1>You made a request</h1> 
	<p>method = %s</p> 
	<p>path = %s</p>
</body>
</html>
`

const applyBodyTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Apply</title>
</head>
<body>
	<h1>"Apply"</h1>
	<form method="post" action="/apply">
		<input type="submit" />
	</form>
</body>
</html>
`

type pageFactory func(string, string) (string, bool)

func main() {
	s := tcp.NewTcpServer(port)
	sh := &SimpleHandler{}
	sh.addFactory(postApplyPageFactory)
	sh.addFactory(getApplyPageFactory)
	sh.addFactory(indexPageFactory)

	s.Add(sh.Handle)

	err := s.Start()
	if err != nil {
		log.Fatalln("Can't start TCP server!", err)
	}
}

type SimpleHandler struct {
	request       string
	method        string
	path          string
	version       string
	pageFactories []pageFactory
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
		var body string
		for _, factory := range sh.pageFactories {
			b, ok := factory(sh.method, sh.path)
			if ok {
				body = b
				break
			}
		}
		resp := fmt.Sprintf(responseTemplate, len(body), body)
		log.Printf("%s\n", sh.request)
		sh.clear()
		return resp, false
	}

	return "", true
}

func (sh *SimpleHandler) addFactory(f pageFactory) {
	sh.pageFactories = append(sh.pageFactories, f)
}

func (sh *SimpleHandler) clear() {
	sh.method = ""
	sh.path = ""
	sh.version = ""
	sh.request = ""
}

func indexPageFactory(method string, path string) (string, bool) {
	if method == "GET" && (path == "/" || path == "/index") {
		return fmt.Sprintf(indexBodyTemplate, method, path), true
	}
	return "", false
}

func getApplyPageFactory(method string, path string) (string, bool) {
	if method == "GET" && path == "/apply" {
		return applyBodyTemplate, true
	}
	return "", false
}

func postApplyPageFactory(method string, path string) (string, bool) {
	if method == "POST" && path == "/apply" {
		return "You did it!", true
	}
	return "", false
}
