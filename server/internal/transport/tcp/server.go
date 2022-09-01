package tcp

import (
	"log"
	"net"
)

type Handler interface {
	OpenNewConnection(conn net.Conn)
}

type Server struct {
	handler Handler
	address string
}

func NewServer(handler Handler, address string) *Server {
	return &Server{
		handler: handler,
		address: address,
	}
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := listen.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go s.handler.OpenNewConnection(conn)
	}
}
