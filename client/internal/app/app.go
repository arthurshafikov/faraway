package app

import (
	"github.com/arthurshafikov/faraway/client/internal/client"
	"github.com/arthurshafikov/faraway/client/internal/services"
)

func Run() {
	services := services.NewServices()

	client.NewClient(services).MakeRequest()
}
