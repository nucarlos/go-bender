package repositories

import (
	"github.com/kamva/mgm/v3"
	"github.com/nucarlos/go-bender/models"
	"go.mongodb.org/mongo-driver/bson"
)

type beerRepository struct {
}

// NewBeerRepository method
func NewBeerRepository() BeerRepository {
	return &beerRepository{}
}

// List :: BeerRepository method
func (r *beerRepository) List() ([]models.BeerItem, error) {
	collection := mgm.Coll(&models.Beer{})
	beers := []models.BeerItem{}

	if err := collection.SimpleFind(&beers, bson.D{}); err != nil {
		return nil, err
	}

	return beers, nil
}

// Create :: BeerRepository method
func (r *beerRepository) Create(payload models.BeerItem) error {
	collection := mgm.Coll(&models.Beer{})
	beer := models.NewBeer(payload.ID, payload.Name, payload.Brewery, payload.Country, payload.Currency, payload.Price)

	if err := collection.Create(beer); err != nil {
		return err
	}

	return nil
}

// Get :: BeerRepository method
func (r *beerRepository) Get(id int64) (*models.BeerItem, error) {
	collection := mgm.Coll(&models.Beer{})
	beer := &models.Beer{}

	if err := collection.First(bson.M{"id": id}, beer); err != nil {
		return nil, err
	}

	beerItem := &models.BeerItem{
		ID:       beer.ID,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Country:  beer.Country,
		Price:    beer.Price,
		Currency: beer.Currency,
	}

	return beerItem, nil
}
