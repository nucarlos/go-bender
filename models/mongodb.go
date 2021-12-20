package models

import (
	"github.com/kamva/mgm/v3"
)

// Beer model mongodb
type Beer struct {
	mgm.DefaultModel `bson:",inline"`
	ID               int64   `json:"id" bson:"id" example:"1"`
	Name             string  `json:"name" bson:"name" example:"Golden"`
	Brewery          string  `json:"brewery" bson:"brewery" example:"Kross"`
	Country          string  `json:"country" bson:"country" example:"Chile"`
	Price            float64 `json:"price" bson:"price" example:"10.5"`
	Currency         string  `json:"currency" bson:"currency" example:"EUR"`
}
