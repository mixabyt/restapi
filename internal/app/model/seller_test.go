package model_test

import (
	"mdl/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeller_Validate(t *testing.T) {
	TestCases := []struct {
		name    string
		s       func() *model.Seller
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Seller {
				return model.TestSeller(t)
			},
			isValid: true,
		},
		{
			name: "empty phone",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.PhoneNumber = ""
				return s
			},
			isValid: false,
		},
		{
			name: "incorect len phone",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.PhoneNumber = "+380"
				return s
			},
			isValid: false,
		},
		{
			name: "invalid phone",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.PhoneNumber = "+380h00000000"
				return s
			},
			isValid: false,
		},
		{
			name: "empty password",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.Password = ""
				return s
			},
			isValid: false,
		},
		{
			name: "empty fisrt name",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.FirstName = ""
				return s
			},
			isValid: false,
		},
		{
			name: "empty second name",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.SecondName = ""
				return s
			},
			isValid: false,
		},
		{
			name: "symbols in fist name",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.FirstName = "dima1"
				return s
			},
			isValid: false,
		},
		{
			name: "symbols in second name",
			s: func() *model.Seller {
				s := model.TestSeller(t)
				s.SecondName = "dima1"
				return s
			},
			isValid: false,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}

}
