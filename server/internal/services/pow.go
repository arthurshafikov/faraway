package services

import (
	"bytes"
	"math/big"
)

type ProofOfWorkCheckerService struct {
	hash       Hash
	target     *big.Int
	difficulty int
}

func NewProofOfWorkCheckerService(hash Hash, difficulty int) *ProofOfWorkCheckerService {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	return &ProofOfWorkCheckerService{
		hash:       hash,
		target:     target,
		difficulty: difficulty,
	}
}

func (pow *ProofOfWorkCheckerService) CheckNonce(hash, nonce []byte) bool {
	result := pow.hash.GetHash(bytes.Join([][]byte{hash, nonce}, []byte{}))

	var intHash big.Int

	return intHash.SetBytes(result).Cmp(pow.target) == -1
}

func (pow *ProofOfWorkCheckerService) GetDifficulty() int {
	return pow.difficulty
}
