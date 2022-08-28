package services

type Hash interface {
	Hash(data []byte) []byte
}

type ProofOfWork interface {
	FindNonce(data []byte, difficulty int) int
}

type Services struct {
	Hash
	ProofOfWork
}

func NewServices() *Services {
	hashService := NewHashService()
	powService := NewProofOfWorkService(hashService)

	return &Services{
		Hash:        hashService,
		ProofOfWork: powService,
	}
}
