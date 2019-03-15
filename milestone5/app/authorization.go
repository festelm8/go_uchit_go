package app

import (
    "encoding/json"
    "fmt"
    "net/http"
    _ "github.com/go-sql-driver/mysql"

    "go_uchit_go/milestone5/model"
    "go_uchit_go/milestone5/utils"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var authData model.AuthData
    err := utils.ParseReqData(r, &authData)
    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: err.Error()})
        return
    }

    user, err := app.DB.GetUserByPhone(authData.Phone)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: "no such user"})
        return
    }

    if !utils.VerifyPassword(user.Pass, authData.Pass) {
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


func (app *App) RegUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var regData model.RegData
    err := utils.ParseReqData(r, &regData)
    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: err.Error()})
        return
    }

    user, _ := app.DB.GetUserByPhone(regData.Phone)
    if user != nil {
        w.WriteHeader(http.StatusForbidden)
        _ = json.NewEncoder(w).Encode(utils.ErrorResponse{Msg: "user exist"})
        return
    }

    regData.Pass, err = utils.HashPassword(regData.Pass)
	if err != nil {
       fmt.Println("failed to create user hash pwd")
       return
    }

    err = app.DB.NewUser(regData)
    if err != nil {
        fmt.Println("failed to create user")
        return
    }

    _ = json.NewEncoder(w)
}