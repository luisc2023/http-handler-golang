package controller

import (
	"html/template"
	"net/http"
)

func CreateOrders(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/createorders.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/getorder.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/listorders.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/updateorder.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/deleteorder.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
