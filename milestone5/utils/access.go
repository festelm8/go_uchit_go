package utils

import (
   "time"
   "github.com/dgrijalva/jwt-go"
   "golang.org/x/crypto/bcrypt"
)


type UserJWTClaims struct {
    UID int64           `json:"uid"`
    jwt.StandardClaims
}

func GenerateJWT(SignKey string, UID int64) (string, error) {
   claims := UserJWTClaims{
       UID,
       jwt.StandardClaims{
           ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
       },
   }
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   tokenString, err := token.SignedString([]byte(SignKey))

   return tokenString, err
}


func HashPassword(pwd string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 15)
    if err != nil {
        return ``, err
    }

    return string(hash), nil
}

func VerifyPassword(hash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}





