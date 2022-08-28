package tcp

import (
	"bufio"
	"bytes"
	"fmt"
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

func (h *Handler) openNewConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	hash := h.services.Hash.RandomHash()
	if _, err := conn.Write([]byte(fmt.Sprintf("%s %v\n", hash, h.services.ProofOfWorkChecker.GetDifficulty()))); err != nil {
		log.Println(err)
	}

	nonce := h.getNonceFromTheResponse(conn)

	if h.services.ProofOfWorkChecker.CheckNonce([]byte(hash), nonce) {
		if _, err := conn.Write(h.services.Quotes.GetQuote()); err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) getNonceFromTheResponse(conn net.Conn) []byte {
	message, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println(err)
	}

	return bytes.TrimRight(message, "\n")
}
