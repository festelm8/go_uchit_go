package model

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
    *sqlx.DB
}

type Datastore interface {
    GetUserByPhone(authData AuthData) (*User, error)
}