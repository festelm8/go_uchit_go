package main

import (
	//"log"
	//"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"

	"go_uchit_go/milestone5/app"
	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/utils"
	"go_uchit_go/milestone5/app/handler"
)


func main() {
	// load config
	config, err := conf.NewConfig("config.yaml").Load()
	utils.CheckError(err)

	app := &app.App{}
	app.Initialize(config)

	// setRouters sets the all required routers
	app.Router.Use(middleware.Logger)
	app.Router.Method("POST","/login", handler.AuthLogin(app))

	app.Run(config)
}
