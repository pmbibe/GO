package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var File string = "DB/authenticate.json"

type User struct {
	User string `json:"user"`
	Pass string `json:"password"`
}
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func readFile(file string) User {
	var user User
	data, _ := os.Open(file)
	d := json.NewDecoder(data)
	d.Decode(&user)
	return user
}

func static(files http.Handler) http.Handler {
	return http.StripPrefix("/static/", files)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/login.html", 301)
}
func login(w http.ResponseWriter, r *http.Request) {
	var session Session
	r.ParseForm()
	// user := r.PostFormValue("uname")
	pass := r.PostFormValue("psw")
	if pass == readFile(File).Pass {
		session.UUID = "123456789"
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		fmt.Fprintln(w, "Đăng nhập thành công")
	} else {
		fmt.Fprintln(w, "Đăng nhập thất bại")
	}
}
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("template"))
	mux.Handle("/static/", static(files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
