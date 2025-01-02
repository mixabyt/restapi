package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	MeasudeUnitsID int    `json:"-"`
	CategoryID     int    `json:"-"`
}

func (p *Product) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.Match(regexp.MustCompile(`^[А-Яа-яіїєґ]+$`))),
	)
}
