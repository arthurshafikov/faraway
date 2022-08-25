package tcp

import (
	"log"
	"net"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handleIncomingRequest(conn net.Conn) {
	response := "123" // todo get from service
	if _, err := conn.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
	if err := conn.Close(); err != nil {
		log.Fatalln(err)
	}
}
