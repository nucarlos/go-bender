package services

import (
	"errors"

	"github.com/nucarlos/go-bender/models"
	"github.com/nucarlos/go-bender/repositories"
)

type beerService struct {
	Repository repositories.BeerRepository
}

// NewBeerService method
func NewBeerService() BeerService {
	x := repositories.NewBeerRepository()
	return &beerService{
		Repository: x,
	}
}

// SearchBeers :: BeerService method
func (s *beerService) SearchBeers() ([]models.BeerItem, error) {
	beers, err := s.Repository.List()

	if err != nil {
		return nil, err
	}

	return beers, nil
}

// AddBeers :: BeerService method
func (s *beerService) AddBeers(id int64, name, brewery, country, currency string, price float64) error {
	exists, _ := s.Repository.Get(id)

	if exists != nil {
		return errors.New("beer already exists")
	}

	payload := models.BeerItem{
		ID:       id,
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}

	if err := s.Repository.Create(payload); err != nil {
		return err
	}

	return nil
}

// SearchBeerById :: BeerService method
func (s *beerService) SearchBeerById(id int64) (*models.BeerItem, error) {
	beer, err := s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return beer, nil
}

// BoxBeerPriceById :: BeerService method
func (s *beerService) BoxBeerPriceById(id int64, currency string, quantiy int64) (*models.BeerBox, error) {
	beer, err := s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	price := beer.Price

	if beer.Currency != currency {
		exchange_rate, err := s.getCurrencyExchanges(*beer, currency, 1)
		if err != nil {
			return nil, err
		}

		price = *exchange_rate
	}

	priceTotal := price * float64(quantiy)

	beerBox := models.BeerBox{
		PriceTotal: priceTotal,
	}

	return &beerBox, nil
}
