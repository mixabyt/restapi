package store

type Store interface {
	Seller() SellerRepository
	Category() CategoryRepository
	MeasureUnits() MeasureUnitsRepository
	Product() ProductRepository
}
