package teststore

import "mdl/internal/app/model"

type SellerRepository struct {
	store   *Store
	sellers map[int]*model.Seller
}

func (s *SellerRepository) Create(seller *model.Seller) error {
	if err := seller.Validate(); err != nil {
		return err
	}
	if err := seller.BeforeCreate(); err != nil {
		return err
	}
	seller.ID = len(s.sellers)
	s.sellers[seller.ID] = seller
	return nil
}

func (s *SellerRepository) GetAll(adminID int) ([]model.Seller, error) {
	sellers := make([]model.Seller, 0, len(s.sellers))
	for _, i := range s.sellers {
		sellers = append(sellers, *i)
	}
	return sellers, nil
}
