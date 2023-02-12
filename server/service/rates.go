package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Tickers map[string]float64

type RatesService struct {
	tickers Tickers
}

func NewRatesService() *RatesService {
	ratesService := new(RatesService)

	ratesService.init()

	return ratesService
}

func (s *RatesService) GetTickers(pairs []string) (Tickers, error) {
	tickers := make(Tickers)

	for _, pair := range pairs {
		symbol := strings.Replace(pair, "-", "", 1)
		price, isFound := s.tickers[symbol]

		if isFound == false {
			msg := fmt.Sprintf("Pair '%s' not found", pair)
			return nil, errors.New(msg)
		}

		tickers[pair] = price
	}

	return tickers, nil
}

func (s *RatesService) init() {
	s.updateTickers()

	go func() {
		for range time.Tick(10 * time.Second) {
			s.updateTickers()
		}
	}()
}

func (s *RatesService) updateTickers() {
	tickers := make(Tickers)
	binanceTickers := fetchBinanceTickers()

	for _, ticker := range binanceTickers {
		price, err := strconv.ParseFloat(ticker.Price, 64)

		if err != nil {
			log.Println(err)
			continue
		}

		tickers[ticker.Symbol] = price
	}

	s.tickers = tickers
}

type binanceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func fetchBinanceTickers() []binanceTicker {
	apiUrl := "https://api.binance.com/api/v3/ticker/price"

	response, err := http.Get(apiUrl)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var binanceTickers []binanceTicker

	if err = json.Unmarshal(body, &binanceTickers); err != nil {
		log.Fatalln(err)
	}

	return binanceTickers
}
