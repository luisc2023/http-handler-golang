package controller

import (
	"html/template"
	"net/http"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/adminhome.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/adminregister.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/adminlogin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/adminupdate.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func AdminDelete(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome Home Page!")
	tmpl, err := template.ParseFiles("templates/admindelete.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
