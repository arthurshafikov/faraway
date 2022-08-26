package services

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
)

type QuoteService struct {
	quoutes [][]byte
}

func NewQuoteService(quotesFilePath string) *QuoteService {
	fileContents, err := ioutil.ReadFile(quotesFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	return &QuoteService{
		quoutes: bytes.Split(fileContents, []byte("\n")),
	}
}

func (s *QuoteService) GetQuote() []byte {
	return s.quoutes[rand.Intn(len(s.quoutes))]
}
