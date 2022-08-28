package services

type Services struct {
	Quotes
	Hash
}

type Quotes interface {
	GetQuote() []byte
}

type Hash interface {
	GetHash(data []byte) []byte
	RandomHash() string
}

type Dependencies struct {
	QuotesFilePath string
}

func NewServices(deps *Dependencies) *Services {
	quoteService := NewQuoteService(deps.QuotesFilePath)
	hash := NewHashService()

	return &Services{
		Quotes: quoteService,
		Hash:   hash,
	}
}
