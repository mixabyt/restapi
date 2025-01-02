package sqlstore

import (
	"mdl/internal/app/model"

	"github.com/lib/pq"
)

type SellerRepository struct {
	store *Store
}

func (s *SellerRepository) Create(seller *model.Seller) error {
	if err := seller.Validate(); err != nil {
		return err
	}
	if err := seller.BeforeCreate(); err != nil {
		return err
	}

	err := s.store.db.QueryRow(
		"INSERT INTO sellers (phone_number, encrypted_password, first_name, second_name, admin_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		seller.PhoneNumber,
		seller.EncryptedPassword,
		seller.FirstName,
		seller.SecondName,
		seller.AdminID,
	).Scan(&seller.ID)
	pgErr, ok := err.(*pq.Error)
	if ok {
		if pgErr.Code == "23505" {
			return errSellerAlreadyExist
		}
	}
	return err
}

func (s *SellerRepository) GetAll(adminID int) ([]*model.Seller, error) {

	rows, err := s.store.db.Query("SELECT id, first_name, second_name, phone_number FROM sellers WHERE admin_id = $1", adminID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var sellers []*model.Seller
	for rows.Next() {
		var seller model.Seller
		if err := rows.Scan(&seller.ID, &seller.FirstName, &seller.SecondName, &seller.PhoneNumber); err != nil {
			return nil, err
		}
		sellers = append(sellers, &seller)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return sellers, nil
}
