package sqlstore

import (
	"database/sql"
	"mdl/internal/app/store"
)

type Store struct {
	db                     *sql.DB
	sellerRepository       *SellerRepository
	categoryRepository     *CategoryRepository
	measureUnitsRepository *MeasureUnitsRepository
	productRepository      *ProductRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Seller() store.SellerRepository {
	if s.sellerRepository != nil {
		return s.sellerRepository
	}
	s.sellerRepository = &SellerRepository{
		store: s,
	}
	return s.sellerRepository
}

func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &CategoryRepository{
		store: s,
	}
	return s.categoryRepository
}

func (s *Store) MeasureUnits() store.MeasureUnitsRepository {
	if s.measureUnitsRepository != nil {
		return s.measureUnitsRepository
	}
	s.measureUnitsRepository = &MeasureUnitsRepository{
		store: s,
	}
	return s.measureUnitsRepository
}

func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}
	s.productRepository = &ProductRepository{
		store: s,
	}
	return s.productRepository
}
