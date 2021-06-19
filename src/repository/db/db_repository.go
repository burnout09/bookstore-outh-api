package db

import (
	"github.com/burnout09/bookstore-outh-api/src/clients/cassandra"
	"github.com/burnout09/bookstore-outh-api/src/domain/access_token"
	"github.com/burnout09/bookstore-outh-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}
