package services

import (
	"crypto/sha256"
)

type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

func (h *HashService) Hash(data []byte) []byte {
	sum256 := sha256.Sum256(data)

	return sum256[:]
}
