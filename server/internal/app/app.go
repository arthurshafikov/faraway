package app

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/arthurshafikov/faraway/server/internal/services"
	"github.com/arthurshafikov/faraway/server/internal/transport/tcp"
	"github.com/arthurshafikov/faraway/server/internal/transport/tcp/handler"
	"golang.org/x/sync/errgroup"
)

var quotesFilePath string

func init() {
	flag.StringVar(&quotesFilePath, "quotesFilePath", "./assets/words-of-wisdom.txt", "Path to the quotes file")
}

func Run() {
	flag.Parse()

	rand.Seed(time.Now().Unix()) // to get different quotes

	services := services.NewServices(&services.Dependencies{
		QuotesFilePath:        quotesFilePath,
		ProofOfWorkDifficulty: 15,
	})

	g, gCtx := errgroup.WithContext(context.Background())
	handler := handler.NewHandler(services)
	tcp.NewServer(handler, ":8090").Run(g, gCtx)

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
