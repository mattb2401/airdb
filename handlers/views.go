package handlers

import (
	"airdb/models"
	"html/template"
	"net/http"
	"os"

	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
)

func Dashboard(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	pwd, _ := os.Getwd()
	temp := pwd + "/ui/home.html"
	t, _ := template.ParseFiles(temp)
	v.SData = make(map[string]interface{})
	v.FlashCount = 0
	ss := scs.Load(r)
	dbs := models.Dbs{}
	userId, _ := ss.GetInt("userId")
	if userId <= 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	v.SData["userId"] = userId
	db.Where("userId = ?", userId).Find(&dbs)
	v.Data = &dbs
	t.Execute(w, v)
}
