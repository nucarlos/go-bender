package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// BeerItemPayload model
type BeerItemPayload struct {
	ID       int64   `json:"id" bson:"id" example:"1"`
	Name     string  `json:"name" bson:"name" example:"Golden"`
	Brewery  string  `json:"brewery" bson:"brewery" example:"Kross"`
	Country  string  `json:"country" bson:"country" example:"Chile"`
	Price    float64 `json:"price" bson:"price" example:"10.5"`
	Currency string  `json:"currency" bson:"currency" example:"EUR"`
}

// Validate method
func (m BeerItemPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
		validation.Field(&m.Name, validation.Required),
		validation.Field(&m.Brewery, validation.Required),
		validation.Field(&m.Country, validation.Required),
		validation.Field(&m.Price, validation.Required),
		validation.Field(&m.Currency, validation.Required, validation.Length(3, 3), validation.By(MyIsISO4217), is.UpperCase, is.Alpha),
	)
}

// BeerBoxPayload model
type BeerBoxPayload struct {
	Currency string `json:"currency" bson:"currency" example:"EUR"`
	Quantity int64  `json:"quantity,omitempty" bson:"quantity" default:"6"`
}

// Validate method
func (m BeerBoxPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Currency, validation.Required, validation.Length(3, 3), validation.By(MyIsISO4217), is.UpperCase, is.Alpha),
		validation.Field(&m.Quantity, validation.Min(1)),
	)
}
