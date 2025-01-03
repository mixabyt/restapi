package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	MeasureUnitsID int    `json:"measure_units_id"`
	CategoryID     int    `json:"-"`
}

func (p *Product) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.Match(regexp.MustCompile(`^[а-яА-Яa-zA-ZіІїЇєЄґҐ0-9\s.,!?;:'"-()]*$`))),
		validation.Field(&p.Price, validation.Required),
		validation.Field(&p.MeasureUnitsID, validation.Required),
		validation.Field(&p.CategoryID, validation.Required),
	)
}
