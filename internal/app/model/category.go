package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Category struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AdminID int    `json:"-"`
}

func (c *Category) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required, validation.Match(regexp.MustCompile(`^[А-Яа-яіїєґ]+$`))),
	)
}
