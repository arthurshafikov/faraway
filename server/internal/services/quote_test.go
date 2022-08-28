package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	quotes         = []byte("someQuote1\nsomeQuote2\nsomeQuote3")
	quotesFilePath = "quotes.txt"
	expectedQuotes = []string{"someQuote1", "someQuote2", "someQuote3"}
)

func TestGetQuote(t *testing.T) {
	createFakeQuotesFile(t)
	quoteService := NewQuoteService(quotesFilePath)

	var resultQuotes [][]byte
	for i := 0; i < 10; i++ {
		resultQuotes = append(resultQuotes, quoteService.GetQuote())
	}

	for _, resultQuote := range resultQuotes {
		require.Contains(t, expectedQuotes, string(resultQuote))
	}
	deleteFakeQuotesFile(t)
}

func createFakeQuotesFile(t *testing.T) {
	t.Helper()
	if err := os.WriteFile(quotesFilePath, quotes, 0600); err != nil { //nolint:gofumpt
		t.Fatal(err)
	}
}

func deleteFakeQuotesFile(t *testing.T) {
	t.Helper()
	if err := os.Remove(quotesFilePath); err != nil {
		t.Fatal(err)
	}
}
