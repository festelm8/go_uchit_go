package main

import (
	//"log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"

	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/utils"
	"go_uchit_go/milestone5/app"
)


func main() {
	config, err := conf.NewConfig("config.yaml").Load()
	utils.CheckError(err)

	instance := &app.App{}
	instance.Initialize(config)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/login", instance.AuthLogin)

	instance.Run(router)
}
