package handlers

import (
	"airdb/helpers"
	"airdb/models"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs"
	"github.com/jinzhu/gorm"
)

type ViewData struct {
	Flash      map[string]string
	FlashCount int
	Data       interface{}
	SData      map[string]interface{}
}

var v ViewData

func SignIn(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	t, _ := template.ParseFiles(pwd + "/ui/login.html")
	v.FlashCount = 0
	t.Execute(w, v)
}

func Authenticate(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	err := r.ParseForm()
	v.FlashCount = 1
	v.Flash = make(map[string]string)
	ss := scs.Load(r)
	if err != nil {
		v.Flash["error"] = "Oh Snap!!. Something happened. Error: " + err.Error()
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if r.FormValue("username") != "" && r.FormValue("password") != "" {
		user := models.User{}
		if db.Where("username = ?", r.FormValue("username")).First(&user).RecordNotFound() {
			v.Flash["error"] = "Oh Snap!!. Wrong username or password"
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			if helpers.CheckPasswordHash(r.FormValue("password"), user.Password) {
				InitFlash()
				ss.PutString(w, "name", user.Name)
				ss.PutString(w, "username", user.Username)
				ss.PutBool(w, "isLoggedIn", true)
				ss.PutString(w, "roles", user.Roles)
				ss.PutInt(w, "userId", user.Id)
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				v.Flash["error"] = "Oh Snap!!. Wrong username or password"
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	} else {
		v.Flash["error"] = "Oh Snap!!. Wrong username or password"
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func Logout(db *gorm.DB, w http.ResponseWriter, r *http.Request, scs *scs.Manager) {
	ss := scs.Load(r)
	err := ss.Clear(w)
	if err != nil {
		v.Flash["error"] = "Oh Snap!!. Error logging out"
		http.Redirect(w, r, "/", http.StatusFound)
	}
	v.Flash = make(map[string]string)
	v.Flash["error"] = "Great !!. Logged out successfully."
	http.Redirect(w, r, "/login", http.StatusFound)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	var p Params
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&p)
	if err != nil {
		helpers.RenderJSON(map[string]string{"code": "109", "error": "Invalid json"}, w)
		return
	}
	pass, _ := helpers.HashPassword(p.Password)
	user := models.User{
		Name:      p.Name,
		Username:  p.Username,
		Password:  pass,
		Status:    "Active",
		Roles:     p.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&user).Error; err != nil {
		helpers.RenderJSON(map[string]string{"code": "109", "error": err.Error()}, w)
		return
	}
	helpers.RenderJSON(map[string]string{"code": "100"}, w)
	return
}

func AllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Find(&users)
	x := make(map[string]interface{})
	x["code"] = "100"
	x["users"] = &users
	helpers.RenderJSON(&x, w)
	return
}

func InitFlash() {
	if v.FlashCount == 0 {
		v.Data = nil
		v.Flash = nil
		v.SData = nil
	}
}
