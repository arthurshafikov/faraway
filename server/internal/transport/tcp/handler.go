package tcp

import (
	"log"
	"net"

	"github.com/arthurshafikov/faraway/server/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) handleIncomingRequest(conn net.Conn) {
	if _, err := conn.Write(h.services.Quotes.GetQuote()); err != nil {
		log.Fatalln(err)
	}
	if err := conn.Close(); err != nil {
		log.Fatalln(err)
	}
}
