package teststore

import "mdl/internal/app/store"

type Store struct {
	sellerRepository *SellerRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.SellerRepository {
	if s.sellerRepository != nil {
		return s.sellerRepository
	}
	s.sellerRepository = &SellerRepository{
		store: s,
	}

	return s.sellerRepository

}
