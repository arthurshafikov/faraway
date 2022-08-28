package services

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

type HashService struct {
}

func NewHashService() *HashService {
	return &HashService{}
}

func (h *HashService) GetHash(data []byte) []byte {
	sum256 := sha256.Sum256(data)

	return sum256[:]
}

func (h *HashService) RandomHash() string {
	randomBytes := make([]byte, 10)
	rand.Read(randomBytes)

	return fmt.Sprintf("%x", h.GetHash(randomBytes))
}
