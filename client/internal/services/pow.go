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
			// I would've removed these fmts in the real application, but there I left them just for you
			fmt.Printf("Nonce = %v\n", nonce)
			fmt.Printf("Result = %s\n", intHash.String())
			fmt.Printf("Target = %s\n", target.String())
			// You can make sure that the result is less than the target
			return nonce
		}
	}

	return 0
}
