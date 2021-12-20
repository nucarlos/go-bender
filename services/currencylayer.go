package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/nucarlos/go-bender/models"
	"github.com/sendgrid/rest"
)

// Response from API CurrencyLayer
type Response struct {
	Success bool               `json:"success"`
	Source  string             `json:"source"`
	Quotes  map[string]float64 `json:"quotes"`
	Error   struct {
		Code int64  `json:"code,omitempty"`
		Info string `json:"info,omitempty"`
	} `json:"error,omitempty"`
}

// getCurrencyExchanges :: BeerService method
func (s *beerService) getCurrencyExchanges(beer models.BeerItem, currency string, quantity int64) (*float64, error) {
	var exchange float64

	// Get URL CurrencyLayer
	baseURL := os.Getenv("CURRENCYLAYER_URL")
	if len(baseURL) == 0 {
		baseURL = "http://api.currencylayer.com"
	}

	// Get Access Key CurrencyLayer
	access_key := os.Getenv("CURRENCYLAYER_ACCESS_KEY")
	if len(access_key) == 0 {
		access_key = "c7e51bed6b2b04ad1ddf6f5135501205"
	}

	// Set query params request
	queryParams := make(map[string]string)
	queryParams["access_key"] = access_key
	queryParams["source"] = beer.Currency
	queryParams["currencies"] = currency
	queryParams["format"] = "1"

	// Send request
	request := rest.Request{
		Method:      rest.Get,
		BaseURL:     baseURL + "/live",
		QueryParams: queryParams,
	}
	resp, err := rest.Send(request)
	if err != nil {
		return nil, errors.New("currencylayer conversion failed")
	}

	// Process response request
	var result Response
	err = json.Unmarshal([]byte(resp.Body), &result)
	if err != nil {
		return nil, errors.New("currencylayer conversion failed")
	}

	if result.Success {
		if v, found := result.Quotes[beer.Currency+currency]; found {
			exchange = v
		}
	} else {
		return nil, errors.New(result.Error.Info)
	}

	// Return exchange rate for currencies
	return &exchange, nil
}
