package store

type Store interface {
	Seller() SellerRepository
	Category() CategoryRepository
}
