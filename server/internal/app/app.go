package app

import "github.com/arthurshafikov/faraway/server/internal/transport/tcp"

func Run() {
	handler := tcp.NewHandler()
	tcp.NewTcpServer(handler, ":8090").Run()
}
