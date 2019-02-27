package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"go_uchit_go/milestone5/utils"
	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/handler"
)

type App struct {
	Router *chi.Mux
	DB     *sqlx.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *conf.Config) {
	// mysql connection string
	mysqlBind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DB.UserName,
		config.DB.UserPassword,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	// http server address and port
	hostBind := fmt.Sprintf("%s:%s",
		config.Host.IP,
		config.Host.Port,
	)

	// open connection to database
	db, err := sqlx.MustConnect("mysql", mysqlBind)
	utils.CheckError(err)
	db = db.Unsafe()
	defer db.Close()

	a.DB = db
	a.Router = chi.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	a.Router.Use(middleware.Logger)

    a.Router.Post("/login", handler.authLogin)

    a.Router.Route("/books", func(r chi.Router) {
        a.Router.Use(UserCtx)
        a.Router.Get("/", getBooks)
    })
}
