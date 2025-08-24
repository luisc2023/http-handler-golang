package controller

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func RegisterUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello Register User!")
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LoginUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello Register User!")
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/update.html") // Your HTML root template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/delete.html") // Your HTML root template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DashboardUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, Home Page...!")
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LogoutUsers(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
