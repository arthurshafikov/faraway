package services

import (
	"bytes"
	"fmt"
	"testing"

	mock_services "github.com/arthurshafikov/faraway/client/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFindNonce(t *testing.T) {
	ctrl := gomock.NewController(t)
	hashMock := mock_services.NewMockHash(ctrl)
	powService := NewProofOfWorkService(hashMock)
	data := []byte("someData")
	hashRequest := bytes.Join([][]byte{data, []byte(fmt.Sprintf("%v", 0))}, []byte{})
	hashResponse := []byte("someResponse")
	gomock.InOrder(
		hashMock.EXPECT().Hash(hashRequest).Times(1).Return(hashResponse),
	)

	nonce := powService.FindNonce(data, 0)

	require.Zero(t, nonce)
}
