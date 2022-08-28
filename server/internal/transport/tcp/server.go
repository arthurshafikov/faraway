package tcp

import (
	"log"
	"net"
)

type TcpServer struct {
	handler *Handler
	address string
}

func NewTcpServer(handler *Handler, address string) *TcpServer {
	return &TcpServer{
		handler: handler,
		address: address,
	}
}

func (s *TcpServer) Run() {
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
			log.Fatalln(err)
		}
		go s.handler.openNewConnection(conn)
	}
}
