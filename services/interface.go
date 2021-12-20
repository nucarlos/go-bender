package services

import (
	"github.com/nucarlos/go-bender/models"
)

// BeerService : represent the beer's services
type BeerService interface {
	SearchBeers() ([]models.BeerItem, error)
	AddBeers(id int64, name, brewery, country, currency string, price float64) error
	SearchBeerById(id int64) (*models.BeerItem, error)
	BoxBeerPriceById(id int64, currency string, quantiy int64) (*models.BeerBox, error)
}
