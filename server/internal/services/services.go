package services

type Services struct {
	Quotes
}

type Quotes interface {
	GetQuote() []byte
}

type Dependencies struct {
	QuotesFilePath string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		Quotes: NewQuoteService(deps.QuotesFilePath),
	}
}
