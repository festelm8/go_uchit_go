package main

import (
	"log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"

	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/app"
)


func main() {
	config, err := conf.NewConfig("config.yaml").Load()
	if err != nil {
		log.Fatal(err)
	}

	instance := &app.App{}
	instance.Init(config)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/me", func(router chi.Router) {
		router.Use(instance.UserCtx)
        router.Get("/", instance.UserInfo)
    })
	router.Post("/login", instance.Login)
	router.Post("/reg", instance.RegUser)

	instance.Run(router)
}
