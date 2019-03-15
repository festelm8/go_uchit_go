package model

import (
	//"fmt"
	"go_uchit_go/milestone5/utils"
)

type User struct {
	ID         int64     `db:"id"        json:"id"`
	Name       string    `db:"name"      json:"name"`
	Phone      string    `db:"phone"     json:"phone"`
	Pass       string    `db:"password"`
}


type AuthData struct {
  Phone string   `json:"phone"     db:"phone"`
  Pass string    `json:"password"`
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

func (db *DB) GetUserByPhone(phone string) (*User, error) {
	var user User
	err := db.Get(&user, `SELECT * FROM users WHERE phone = ?`, phone)
	if err != nil {
        return nil, err
    }

    return &user, nil
}


type RegData struct {
  Name   string   `db:"name"      json:"name"`
  Phone  string   `db:"phone"     json:"phone"`
  Pass   string   `db:"password"  json:"password"`
}

func (reqData *RegData) OK() error {
	var errors utils.ErrMissingField
	if len(reqData.Name) == 0 {
		errors.Fields = append(errors.Fields, "name")
	}
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

func (db *DB) NewUser(regData RegData) error {
	pstmt, err := db.PrepareNamed(`INSERT INTO users (name, phone, password) VALUES (:name, :phone, :password)`)
	if err != nil {
        return err
	}

	_, err = pstmt.Exec(regData)
	if err != nil {
        return err
    }

    return nil
}