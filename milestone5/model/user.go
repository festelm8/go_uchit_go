package model

import (
	//"fmt"
	"go_uchit_go/milestone5/utils"
)

type AuthData struct {
  Phone string `json:"phone"`
  Pass string `json:"password"`
}

func (reqData *AuthData) OK() error {
	var errors utils.ErrMissingField
	if len(reqData.Phone) == 0 {
		errors.Fields = append(errors.Fields, "phone")
	}
	if len(reqData.Pass) == 0 {
		errors.Fields = append(errors.Fields, "password")
	}
	if len(errors.Fields) != 0 {
		return errors
	}
	return nil
}

type AuthToken struct {
  Token string `json:"auth_token"`
}

type User struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Phone      string    `db:"phone" json:"phone"`
	Pass       string    `db:"password"`
}

func (db *DB) GetUserByPhone(authData AuthData) (*User, error) {
	var user User
	pstmt, err := db.PrepareNamed(`SELECT * FROM users WHERE phone = :phone`)
	if err != nil {
        return nil, err
	}

	err = pstmt.Get(&user, authData)
	if err != nil {
        return nil, err
    }

    return &user, nil
}
