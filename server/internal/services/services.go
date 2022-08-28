package services

type Services struct {
	Quotes
	Hash
	ProofOfWorkChecker
}

type Quotes interface {
	GetQuote() []byte
}

type Hash interface {
	GetHash(data []byte) []byte
	RandomHash() string
}

type ProofOfWorkChecker interface {
	CheckNonce(hash, nonce []byte) bool
	GetDifficulty() int
}

type Dependencies struct {
	QuotesFilePath        string
	ProofOfWorkDifficulty int
}

func NewServices(deps *Dependencies) *Services {
	quoteService := NewQuoteService(deps.QuotesFilePath)
	hash := NewHashService()
	proofOfWorkChecker := NewProofOfWorkCheckerService(hash, deps.ProofOfWorkDifficulty)

	return &Services{
		Quotes:             quoteService,
		Hash:               hash,
		ProofOfWorkChecker: proofOfWorkChecker,
	}
}
