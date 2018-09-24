package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

func Dashboard(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	t, _ := template.ParseFiles(pwd + "/views/dash.html")
	v.FlashCount = 0
	t.Execute(w, v)
}
