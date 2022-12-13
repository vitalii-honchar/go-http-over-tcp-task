package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
)

type TcpHandler func(any) (any, bool)

type TcpServer struct {
	port     string
	handlers []TcpHandler
}

func NewTcpServer(port int) *TcpServer {
	return &TcpServer{
		port: fmt.Sprintf(":%v", port),
	}
}

func (ts *TcpServer) Add(h TcpHandler) *TcpServer {
	ts.handlers = append(ts.handlers, h)
	return ts
}

func (ts *TcpServer) Start() error {
	log.Printf("Starting TCP server at port %v\n", ts.port)

	ln, err := net.Listen("tcp", ts.port)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		ts.handle(conn)
	}
}

func (ts *TcpServer) handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "I see you connected\n")
}
