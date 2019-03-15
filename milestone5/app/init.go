package app

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/go-chi/chi"


	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/model"
)

type App struct {
	DB        model.Datastore
	Conf      *conf.Config
}

func (app *App) Initialize(config *conf.Config) {
	mysqlBind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DB.UserName,
		config.DB.UserPassword,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	db := sqlx.MustConnect("mysql", mysqlBind)
	db = db.Unsafe()

	app.Conf = config
	app.DB = &model.DB{db}
}

func (app *App) Run(router *chi.Mux) {
	hostBind := fmt.Sprintf("%s:%s",
		app.Conf.Host.IP,
		app.Conf.Host.Port,
	)

	fmt.Println(">> Here we go! Server is run on :5000")
	log.Fatal(http.ListenAndServe(hostBind, router))
}
