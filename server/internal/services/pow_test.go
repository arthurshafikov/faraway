package services

import (
	"fmt"
	"testing"

	mock_services "github.com/arthurshafikov/faraway/server/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCheckNonce(t *testing.T) {
	ctrl := gomock.NewController(t)
	hashServiceMock := mock_services.NewMockHash(ctrl)
	powService := NewProofOfWorkCheckerService(hashServiceMock, 15)
	hash := "sajkjij3124oj12j31jksaf"
	gomock.InOrder(
		hashServiceMock.EXPECT().GetHash([]byte(fmt.Sprintf("%s%v", hash, 15))).Return([]byte("jhfujakwjirhtuw")),
	)

	result := powService.CheckNonce([]byte(hash), []byte("15"))
	require.True(t, result)
}

func TestGetDifficulty(t *testing.T) {
	powService := NewProofOfWorkCheckerService(nil, 15)

	result := powService.GetDifficulty()
	require.Equal(t, 15, result)
}
