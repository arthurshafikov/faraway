package app

import (
	"math/rand"
	"time"

	"github.com/arthurshafikov/faraway/server/internal/services"
	"github.com/arthurshafikov/faraway/server/internal/transport/tcp"
)

func Run() {
	rand.Seed(time.Now().Unix()) // to get different quotes

	services := services.NewServices(&services.Dependencies{
		QuotesFilePath: "./assets/words-of-wisdom.txt", // todo config
	})

	handler := tcp.NewHandler(services)
	tcp.NewTcpServer(handler, ":8090").Run() // todo config
}
