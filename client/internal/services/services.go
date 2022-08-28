package services

type Hash interface {
	Hash(data []byte) []byte
}

type Services struct {
	Hash
}

func NewServices() *Services {
	hashService := NewHashService()

	return &Services{
		Hash: hashService,
	}
}
