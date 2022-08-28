package services

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	hashService := NewHashService()
	data := []byte("Some")
	expected := sha256.Sum256(data)

	result := hashService.Hash(data)

	require.Equal(t, expected[:], result)
}
