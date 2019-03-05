package handler

import (
    "encoding/json"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"

    "go_uchit_go/milestone5/app"
    "go_uchit_go/milestone5/utils"
)

type AuthData struct {
  Phone string `json:"phone_number"`
  Pass string `json:"password"`
}

type AuthToken struct {
  Token string `json:"auth_token"`
}

func AuthLogin(a *app.App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        var authData AuthData
        _ = json.NewDecoder(r.Body).Decode(&authData)
        if authData.Pass != "kok" {
            w.WriteHeader(http.StatusForbidden)
            json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: "wrong password"})
            return
        }

        validToken, err := utils.GenerateJWT()
        if err != nil {
            fmt.Println("Failed to generate token")
            return
        }
        json.NewEncoder(w).Encode(AuthToken{Token: validToken})
    })
}
