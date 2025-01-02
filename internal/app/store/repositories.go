package store

import "mdl/internal/app/model"

type SellerRepository interface {
	Create(*model.Seller) error
	GetAll(int) ([]*model.Seller, error)
}

type CategoryRepository interface {
	Create(*model.Category) error
	GetAll(int) ([]*model.Category, error)
}
