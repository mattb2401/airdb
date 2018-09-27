package handlers

import (
	"airdb/helpers"
	"airdb/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
)

func RunQuery(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	type Params struct {
		Query  string `json:"query"`
		DbId   int    `json:"dbId"`
		UserId string `json:"userId"`
	}
	var p Params
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&p)
	if err != nil {
		log.Print(err)
		helpers.RenderJSON(map[string]string{"code": "109", "error": "Invalid json"}, w)
		return
	}
	Db := models.Db{}
	if db.Where("id = ?", p.DbId).Where("userId = ?", p.UserId).First(&Db).RecordNotFound() {
		helpers.RenderJSON(map[string]string{"code": "109", "error": "Invalid Db"}, w)
		return
	} else {
		dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", Db.Username, Db.Password, Db.Host, Db.Port, Db.Name)
		dbCon, err := sql.Open(Db.DBSchema, dbURI)
		if err != nil {
			helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
			return
		}
		rows, err := dbCon.Query(p.Query)
		if err != nil {
			helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
			return
		}
		cols, err := rows.Columns()
		if err != nil {
			helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
			return
		}
		var vals []map[string]interface{}
		for rows.Next() {
			columns, err := rows.ColumnTypes()
			if err != nil {
				helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
				return
			}
			values := make([]interface{}, len(columns))
			object := map[string]interface{}{}
			for i, column := range columns {
				object[column.Name()] = new(*string)
				values[i] = object[column.Name()]
			}

			err = rows.Scan(values...)
			if err != nil {
				helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
				return
			}
			vals = append(vals, object)
		}
		x := make(map[string]interface{})
		x["code"] = "100"
		x["columns"] = cols
		x["data"] = vals
		dbCon.Close()
		helpers.RenderJSON(&x, w)
		return
	}

}
