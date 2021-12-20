package models

// BeerItem model
type BeerItem struct {
	ID       int64   `json:"id" bson:"id" example:"1"`
	Name     string  `json:"name" bson:"name" example:"Golden"`
	Brewery  string  `json:"brewery" bson:"brewery" example:"Kross"`
	Country  string  `json:"country" bson:"country" example:"Chile"`
	Price    float64 `json:"price" bson:"price" example:"10.5"`
	Currency string  `json:"currency" bson:"currency" example:"EUR"`
}

// BeerBox model
type BeerBox struct {
	PriceTotal float64 `json:"price_total" bson:"price_total"`
}

// NewBeer create beer entry
func NewBeer(id int64, name, brewery, country, currency string, price float64) *Beer {
	return &Beer{
		ID:       id,
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}
}
