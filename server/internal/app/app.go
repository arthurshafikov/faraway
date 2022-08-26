package app

import (
	"flag"
	"math/rand"
	"time"

	"github.com/arthurshafikov/faraway/server/internal/services"
	"github.com/arthurshafikov/faraway/server/internal/transport/tcp"
)

var quotesFilePath string

func init() {
	flag.StringVar(&quotesFilePath, "quotesFilePath", "./assets/words-of-wisdom.txt", "Path to the quotes file")
}

func Run() {
	flag.Parse()

	rand.Seed(time.Now().Unix()) // to get different quotes

	services := services.NewServices(&services.Dependencies{
		QuotesFilePath: quotesFilePath,
	})

	handler := tcp.NewHandler(services)
	tcp.NewTcpServer(handler, ":8090").Run() // todo config
}
