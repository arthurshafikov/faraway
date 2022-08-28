package services

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWorkService struct {
	hashService Hash
}

func NewProofOfWorkService(hashService Hash) *ProofOfWorkService {
	return &ProofOfWorkService{
		hashService: hashService,
	}
}

func (pow *ProofOfWorkService) FindNonce(data []byte, difficulty int) int {
	var intHash big.Int

	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	var hash []byte
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		hash = pow.hashService.Hash(bytes.Join([][]byte{data, []byte(fmt.Sprintf("%v", nonce))}, []byte{}))

		intHash.SetBytes(hash)

		if intHash.Cmp(target) == -1 {
			fmt.Printf("Nonce = %v\n", nonce)
			fmt.Printf("Result = %s\n", intHash.String())
			fmt.Printf("Target = %s\n", target.String())
			return nonce
		}
	}

	return 0
}
