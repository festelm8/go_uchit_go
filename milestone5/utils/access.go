package utils

import (
   //"context"
   //"encoding/json"
   //"fmt"
   //"github.com/go-chi/chi"
   //"github.com/go-chi/chi/middleware"
   //"math/rand"
   //"net/http"
   //"strconv"
   "time"
   //_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
   "github.com/dgrijalva/jwt-go"
   "golang.org/x/crypto/bcrypt"
)


var mySigningKey = []byte("cockkekkok")

type UserJWTClaims struct {
    UID int64           `json:"uid"`
    jwt.StandardClaims
}

func GenerateJWT(UID int64) (string, error) {
   claims := UserJWTClaims{
       UID,
       jwt.StandardClaims{
           ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
       },
   }
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   tokenString, err := token.SignedString(mySigningKey)

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




//func UserCtx(next http.Handler) http.Handler {
//    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//        if r.Header["Authorization"] != nil {
//
//            token, err := jwt.ParseWithClaims(r.Header["Authorization"][0], &UserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
//                return mySigningKey, nil
//            })
//
//            if err == nil {
//                if claims, ok := token.Claims.(*UserJWTClaims); ok && token.Valid {
//                    currentUser := User{ID: "123", Name: "CockUser"}
//                    if currentUser.ID == claims.UID {
//                        ctx := context.WithValue(r.Context(), "currentUser", currentUser)
//                        next.ServeHTTP(w, r.WithContext(ctx))
//                        return
//                    }
//                }
//            }
//        }
//
//        ctx := context.WithValue(r.Context(), "currentUser", nil)
//        next.ServeHTTP(w, r.WithContext(ctx))
//    })
//}

