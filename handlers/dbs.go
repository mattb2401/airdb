package handlers

import (
	"airdb/helpers"
	"airdb/models"
	"encoding/json"
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
)

func AddDb(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	type Params struct {
		DBSchema string `json:"dbschema"`
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		UserId   int    `json:"userId"`
	}
	var p Params
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&p)
	if err != nil {
		helpers.RenderJSON(map[string]string{"code": "109", "error": "Invalid json"}, w)
		return
	}
	DB := models.Db{
		Username: p.Username,
		Password: p.Password,
		DBSchema: p.DBSchema,
		Host:     p.Host,
		Port:     p.Port,
		Name:     p.Name,
		UserId:   p.UserId,
	}
	if err := db.Create(&DB).Error; err != nil {
		helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
		return
	} else {
		helpers.RenderJSON(map[string]string{"code": "100"}, w)
		return
	}
}

func AllDbs(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	type Params struct {
		UserId int `json:"userId"`
	}
	var p Params
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&p)
	if err != nil {
		helpers.RenderJSON(map[string]string{"code": "109", "error": "Invalid json"}, w)
		return
	}
	Dbs := models.Dbs{}
	db.Where("userId = ?", p.UserId).Find(&Dbs)
	x := make(map[string]interface{})
	x["code"] = "100"
	x["Dbs"] = Dbs
	helpers.RenderJSON(&x, w)
	return
}
