package db

import (
	"github.com/santoshr1016/bookstore_oauth-api/src/domain/access_token"
	"github.com/santoshr1016/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository  {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string)(*access_token.AccessToken, *errors.RestError)
}
type dbRepository struct {

}

func (db *dbRepository)GetById(id string)(*access_token.AccessToken, *errors.RestError)  {
	return nil, errors.NewInternalServerError("Error in DB")
}
