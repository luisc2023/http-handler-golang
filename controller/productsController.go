package controller

import (
	"html/template"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/createproducts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/getproduct.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/listpropducts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/updateproduct.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/deleteproduct.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
