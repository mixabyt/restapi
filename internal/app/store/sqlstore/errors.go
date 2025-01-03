package sqlstore

import "errors"

var (
	errSellerAlreadyExist  = errors.New("seller already exist")
	errProductForeignKey   = errors.New("such category or measure units does not exist")
	errProductAlreadyExist = errors.New("product already exist")
	errDeniedAccess        = errors.New("denied access")
)
