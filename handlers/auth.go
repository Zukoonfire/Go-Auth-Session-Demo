package handlers

import (
	"log"
	"net/http"
	"text/template"

	"authsession/models"

	"github.com/google/uuid"
)

var templates = template.Must(template.ParseFiles("templates/login.html", "templates/register.html", "templates/welcome.html"))


func RegisterPage(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "register.html", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/register", http.StatusSeeOther)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    err := models.RegisterUser(email, password)
    if err != nil {
        http.Error(w, "Unable to register user", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "login.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
  log.Println("Coming here?")
	if r.Method != http.MethodPost {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    if models.AuthenticateUser(email, password) {
        sessionID := uuid.NewString()
        http.SetCookie(w, &http.Cookie{
            Name:  "session_id",
            Value: sessionID,
            Path:  "/",
        })
        http.Redirect(w, r, "/welcome", http.StatusSeeOther)
    } else {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_id")
    if err != nil || cookie.Value == "" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    templates.ExecuteTemplate(w, "welcome.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:   "session_id",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(w, cookie)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
