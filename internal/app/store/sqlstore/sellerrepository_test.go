package sqlstore_test

import (
	"mdl/internal/app/model"
	"mdl/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSellerRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("sellers")
	s := sqlstore.New(db)
	seller := model.TestSeller(t)
	assert.NoError(t, s.Seller().Create(seller))
	assert.NotNil(t, seller.ID)
}
