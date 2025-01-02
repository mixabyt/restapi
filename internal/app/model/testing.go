package model

import "testing"

func TestSeller(t *testing.T) *Seller {
	t.Helper()

	return &Seller{
		FirstName:   "Ivan",
		SecondName:  "Fomin",
		PhoneNumber: "+380123456789",
		Password:    "password",
	}
}
