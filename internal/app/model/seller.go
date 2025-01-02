package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	ID                int    `json:"id"`
	FirstName         string `json:"first_name"`
	SecondName        string `json:"second_name"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (s *Seller) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^\+?380\d{9}$`))),
		validation.Field(&s.Password, validation.By(requiredIf(s.EncryptedPassword == "")), validation.Length(6, 20)),
		validation.Field(&s.FirstName, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-zА-Яа-яІіЇїЄєҐґ'’\s-]+$`))),
		validation.Field(&s.SecondName, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-zА-Яа-яІіЇїЄєҐґ'’\s-]+$`))),
	)
}

func (s *Seller) BeforeCreate() error {
	if len(s.Password) > 0 {
		enc, err := encryptString(s.Password)
		if err != nil {
			return err
		}
		s.EncryptedPassword = enc
	}
	return nil
}
func (s *Seller) Sanitize() {
	s.Password = ""
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
