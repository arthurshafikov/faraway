package handler

import (
	"bufio"
	"fmt"
	"net"
	"testing"

	"github.com/arthurshafikov/faraway/server/internal/services"
	mock_services "github.com/arthurshafikov/faraway/server/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

var (
	randomHash  = "3321asfa24aga3yhbasdsa"
	difficulty  = 15
	nonce       = 23
	randomQuote = "Some Quote ~ some Author"
)

func TestOpenNewConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	quotesServiceMock := mock_services.NewMockQuotes(ctrl)
	hashServiceMock := mock_services.NewMockHash(ctrl)
	powServiceMock := mock_services.NewMockProofOfWorkChecker(ctrl)
	services := services.Services{
		Quotes:             quotesServiceMock,
		Hash:               hashServiceMock,
		ProofOfWorkChecker: powServiceMock,
	}
	gomock.InOrder(
		hashServiceMock.EXPECT().RandomHash().Return(randomHash),
		powServiceMock.EXPECT().GetDifficulty().Return(difficulty),
		powServiceMock.EXPECT().CheckNonce([]byte(randomHash), []byte(fmt.Sprintf("%v", nonce))).Return(true),
		quotesServiceMock.EXPECT().GetQuote().Return([]byte(randomQuote)),
	)
	handler := NewHandler(&services)
	connIn, connOut := net.Pipe()
	g := &errgroup.Group{}
	g.Go(func() error {
		message, err := bufio.NewReader(connOut).ReadString('\n')
		require.NoError(t, err)
		require.Equal(t, fmt.Sprintf("%s %v\n", randomHash, difficulty), message)

		_, err = fmt.Fprintf(connOut, "%v\n", nonce)
		require.NoError(t, err)

		message, err = bufio.NewReader(connOut).ReadString('\n')
		require.NoError(t, err)
		require.Equal(t, randomQuote+"\n", message)

		return nil
	})

	handler.OpenNewConnection(connIn)

	require.NoError(t, connIn.Close())
	require.NoError(t, connOut.Close())
	require.NoError(t, g.Wait())
}
