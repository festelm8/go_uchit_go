package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"go_uchit_go/milestone5/conf"
)

// App struct
type App struct {
	Router *chi.Mux
	DB     *sqlx.DB
	Conf   *conf.Config
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *conf.Config) {
	mysqlBind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DB.UserName,
		config.DB.UserPassword,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	db := sqlx.MustConnect("mysql", mysqlBind)
	db = db.Unsafe()
	defer db.Close()

	a.Conf = config
	a.DB = db
	a.Router = chi.NewRouter()
}

// Run the app on it's router
func (a *App) Run(config *conf.Config) {
	hostBind := fmt.Sprintf("%s:%s",
		config.Host.IP,
		config.Host.Port,
	)

	fmt.Println(">> Here we go! Server is run on :5000")
	log.Fatal(http.ListenAndServe(hostBind, a.Router))
}
