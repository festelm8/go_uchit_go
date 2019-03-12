package model

import (
    "github.com/jmoiron/sqlx"

    "go_uchit_go/milestone5/utils"
)

type AuthData struct {
  Phone string `json:"phone_number"`
  Pass string `json:"password"`
}

type AuthToken struct {
  Token string `json:"auth_token"`
}

type User struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Phone      string    `db:"phone" json:"phone"`
}

func GetUserByPhone(db *sqlx.DB, authData AuthData) (*User, error) {
	var user User
	pstmt, err := db.PrepareNamed(`SELECT * FROM place WHERE phone = :phone_number`)
	utils.CheckError(err)

	err = pstmt.Select(&user, authData)
	utils.CheckError(err)

    return &user, nil
}

func (us *UserHandler) getUser(id int64) (*utils.ResultTransformer, error) {

	// concurrency safe
	us.lck.RLock()
	defer us.lck.RUnlock()

	user := User{}

	err := us.db.Get(&user, "select * from tbl_users where id = ?", id)
	if err != nil {
		return nil, err
	}

	header := models.Header{Status: "ok", Count: 1, Data: user}
	result := utils.NewResultTransformer(header)

	return result, nil
}