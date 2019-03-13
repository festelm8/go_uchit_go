package model

import (
	//"fmt"
	"go_uchit_go/milestone5/utils"
)

type AuthData struct {
  Phone string `json:"phone"`
  Pass string `json:"password"`
}

func (data *AuthData) OK() error {
	var errors utils.ErrMissingField
	if len(data.Phone) == 0 {
		errors = append(errors, "phone")
	}
	if len(data.Pass) == 0 {
		errors = append(errors, "password")
	}
	return errors
}

type AuthToken struct {
  Token string `json:"auth_token"`
}

type User struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Phone      string    `db:"phone" json:"phone"`
}

func (db *DB) GetUserByPhone(authData AuthData) (*User, error) {
	var user User
	pstmt, err := db.PrepareNamed(`SELECT * FROM place WHERE phone = :phone`)
	if err != nil {
        return nil, err
    }

	err = pstmt.Select(&user, authData)
	if err != nil {
        return nil, err
    }

    return &user, nil
}
