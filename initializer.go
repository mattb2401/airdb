package main

import (
	"airdb/config"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Init struct {
	Route *mux.Router
	DB    *gorm.DB
}

//Initialize application
func (i *Init) Initialize() {
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name, config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		panic(err)
	}
	i.DB = db
	i.Route = mux.NewRouter()
	i.setRoutes()
}

func (i *Init) setRoutes() {

}

// Get requests handler with mux
func (i *Init) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	i.Route.HandleFunc(path, f).Methods("GET")
}

// Post requests handler with mux
func (i *Init) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	i.Route.HandleFunc(path, f).Methods("POST")
}

func (i *Init) startRoutines() {

}

// Run application
func (i *Init) Run(host string) {
	fmt.Println("Running airdb on " + host)
	err := http.ListenAndServe(host, i.Route)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
