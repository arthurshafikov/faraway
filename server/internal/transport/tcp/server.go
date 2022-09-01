package tcp

import (
	"context"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
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

func (s *Server) Run(g *errgroup.Group, gCtx context.Context) {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalln(err)
	}
	g.Go(func() error {
		<-gCtx.Done()

		if err := listen.Close(); err != nil {
			log.Fatalln(err)
		}

		return nil
	})
	for {
		conn, err := listen.Accept()
		if err != nil {
			select {
			case <-gCtx.Done():
				return
			default:
			}
			log.Println(err)
			break
		}
		g.Go(func() error {
			s.handler.OpenNewConnection(conn)

			return nil
		})
	}
}
