package app

import (
    "encoding/json"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"

    "go_uchit_go/milestone5/model"
    "go_uchit_go/milestone5/utils"
)

func (app *App) AuthLogin(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var authData model.AuthData
    err := utils.ParseReqData(r, &authData)
    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: err.Error()})
        return
    }

    user, err := app.DB.GetUserByPhone(authData)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: "no such user"})
        return
    }

    if authData.Pass != user.Pass {
        w.WriteHeader(http.StatusForbidden)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: "wrong password"})
        return
    }

    validToken, err := utils.GenerateJWT(user.ID)
    if err != nil {
        fmt.Println("Failed to generate token")
        return
    }

    _ = json.NewEncoder(w).Encode(model.AuthToken{Token: validToken})
}