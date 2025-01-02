package sqlstore

import "mdl/internal/app/model"

type CategoryRepository struct {
	store *Store
}

func (c *CategoryRepository) Create(category *model.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}
	return c.store.db.QueryRow(
		"INSERT INTO categories (name) VALUES ($1) RETURNING id",
		category.Name,
	).Scan(&category.ID)
}

func (c *CategoryRepository) GetAll(adminID int) ([]*model.Category, error) {
	rows, err := c.store.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
