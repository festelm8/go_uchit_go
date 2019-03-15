package model

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
    *sqlx.DB
}

type Datastore interface {
    GetUserByPhone(phone string) (*User, error)
    GetUserByID(ID int64) (*User, error)
    NewUser(regData RegData) error
}