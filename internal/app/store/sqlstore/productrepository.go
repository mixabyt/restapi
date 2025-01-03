package sqlstore

import (
	"mdl/internal/app/model"

	"github.com/lib/pq"
)

type ProductRepository struct {
	store *Store
}

func (p *ProductRepository) Create(product *model.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	err := p.store.db.QueryRow(
		"INSERT INTO products (name, price, measure_units_id, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		&product.Name,
		&product.Price,
		&product.MeasureUnitsID,
		&product.CategoryID,
	).Scan(&product.ID)

	pgErr, ok := err.(*pq.Error)
	if ok {
		switch pgErr.Code {
		case "23503":
			return errProductForeignKey
		case "23505":
			return errProductAlreadyExist
		}
	}
	return err
}

func (p *ProductRepository) Get(adminID, categoryID int) ([]*model.Product, error) {
	var access bool
	err := p.store.db.QueryRow("SELECT EXISTS (SELECT 1 FROM categories WHERE admin_id = $1 and id = $2) AS access;", adminID, categoryID).Scan(&access)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, errDeniedAccess
	}

	rows, err := p.store.db.Query("SELECT id,name,price,measure_units_id FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.MeasureUnitsID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
