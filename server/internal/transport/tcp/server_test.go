package tcp

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	mock_server "github.com/arthurshafikov/faraway/server/internal/transport/tcp/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

var address = "localhost:1234"

func TestRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ctrl := gomock.NewController(t)
	handlerMock := mock_server.NewMockHandler(ctrl)
	gomock.InOrder(
		handlerMock.EXPECT().OpenNewConnection(gomock.Any()),
	)
	s := NewServer(handlerMock, address)
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		s.Run(g, gCtx)

		return nil
	})
	var (
		conn net.Conn
		err  error
	)
	for {
		select {
		case <-ctx.Done():
			t.Fatal("connection timeout")
		default:
		}
		conn, err = net.Dial("tcp", address)
		if err == nil {
			break
		}
	}
	defer func() {
		require.NoError(t, conn.Close())
	}()

	_, err = fmt.Fprintf(conn, "12")
	require.NoError(t, err)

	cancel()
	require.NoError(t, g.Wait())
}
