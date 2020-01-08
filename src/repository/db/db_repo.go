package db

import (
	"github.com/santoshr1016/bookstore_oauth-api/src/clients/cassandra"
	"github.com/santoshr1016/bookstore_oauth-api/src/domain/access_token"
	"github.com/santoshr1016/bookstore_oauth-api/src/utils/errors"
	"pkg/mod/github.com/gocql/gocql@v0.0.0-20200103014340-68f928edb90a"
)

const(
	queryGetAccessToken="Select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryCreateAccessToken="insert into access_tokens(access_token, user_id, client_id, expires) values (?,?,?,?);"
	queryUpdateAccessToken="update access_tokens set expires=? where access_token=?;"
)

func NewRepository() DbRepository  {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string)(*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken)*errors.RestError
	UpdateExpirationTime(access_token.AccessToken)*errors.RestError
}
type dbRepository struct {

}

func (db *dbRepository)GetById(id string)(*access_token.AccessToken, *errors.RestError) {
	//TODO implement get access token from Cassandra DB
	//session, err := cassandra.GetSession()
	//if err != nil {
	//	//panic(err)
	//	return nil, errors.NewInternalServerError(err.Error())
	//}
	//defer session.Close()
	//fmt.Println("Connection created")
	// TODO Using Single Global Cassandra Session
	var result access_token.AccessToken
	//if err := session.Query(queryGetAccessToken, id).Scan(
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound{
			return nil, errors.NewNotFoundError("No access token found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
	//return nil, errors.NewInternalServerError("Error in DB")
}

func (db *dbRepository)Create(at access_token.AccessToken) *errors.RestError  {
	//session, err := cassandra.GetSession()
	//if err != nil {
	//	//panic(err)
	//	return errors.NewInternalServerError(err.Error())
	//}
	//defer session.Close()
	// TODO Using Single Global Cassandra Session
	//if err := session.Query(queryCreateAccessToken,
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (db *dbRepository)UpdateExpirationTime(at access_token.AccessToken) *errors.RestError  {
	//session, err := cassandra.GetSession()
	//if err != nil {
	//	//panic(err)
	//	return errors.NewInternalServerError(err.Error())
	//}
	//defer session.Close()
	// TODO Using Single Global Cassandra Session
	if err := cassandra.GetSession().Query(queryUpdateAccessToken,
		at.AccessToken,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}