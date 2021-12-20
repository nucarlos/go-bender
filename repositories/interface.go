package repositories

import (
	"github.com/nucarlos/go-bender/models"
)

// BeerRepository : represent the beer's repositories
type BeerRepository interface {
	List() ([]models.BeerItem, error)
	Create(models.BeerItem) error
	Get(id int64) (*models.BeerItem, error)
}
