package main

import (
	"airdb/config"
	"airdb/handlers"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alexedwards/scs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type Init struct {
	Route   *mux.Router
	Session *scs.Manager
	DB      *gorm.DB
}

//Initialize application
func (i *Init) Initialize() {
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name, config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		panic(err)
	}
	i.DB = db
	i.Route = mux.NewRouter()
	i.Route.HandleFunc(`/ui/assets/{path:[a-zA-Z0-9=\-\/\.\_]+}`, use(handleStatic))
	i.Session = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")
	i.setRoutes()
}

func (i *Init) setRoutes() {
	i.Get("/login", i.signIn)
	i.Post("/authenticate", i.authenticate)
	i.Get("/", i.dashboard)
	i.Post("/users/register", i.register)
	i.Get("/users", i.allUsers)
	i.Post("/dbs", i.allDbs)
	i.Post("/dbs/add", i.addDB)
	i.Get("/logout", i.logout)
	i.Post("/queries/run", i.runQuery)
}

// Get requests handler with mux
func (i *Init) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	i.Route.HandleFunc(path, f).Methods("GET")
}

// Post requests handler with mux
func (i *Init) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	i.Route.HandleFunc(path, f).Methods("POST")
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	static := vars["path"]
	http.ServeFile(w, r, filepath.Join("ui/assets", static))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func (i *Init) signIn(w http.ResponseWriter, r *http.Request) {
	handlers.SignIn(w, r)
}

func (i *Init) authenticate(w http.ResponseWriter, r *http.Request) {
	handlers.Authenticate(i.DB, w, r, i.Session)
}

func (i *Init) addDB(w http.ResponseWriter, r *http.Request) {
	handlers.AddDb(i.DB, w, r, i.Session)
}

func (i *Init) allDbs(w http.ResponseWriter, r *http.Request) {
	handlers.AllDbs(i.DB, w, r, i.Session)
}

func (i *Init) runQuery(w http.ResponseWriter, r *http.Request) {
	handlers.RunQuery(i.DB, w, r, i.Session)
}

func (i *Init) register(w http.ResponseWriter, r *http.Request) {
	handlers.CreateUser(i.DB, w, r)
}

func (i *Init) allUsers(w http.ResponseWriter, r *http.Request) {
	handlers.AllUsers(i.DB, w, r)
}

func (i *Init) dashboard(w http.ResponseWriter, r *http.Request) {
	handlers.Dashboard(i.DB, w, r, i.Session)
}

func (i *Init) logout(w http.ResponseWriter, r *http.Request) {
	handlers.Logout(i.DB, w, r, i.Session)
}

// Run application
func (i *Init) Run(host string) {
	fmt.Println("Running airdb on " + host)
	err := http.ListenAndServe(host, i.Route)
	if err != nil {
		fmt.Println("Error occurred while starting web server with error " + err.Error())
		os.Exit(103)
	}
}
