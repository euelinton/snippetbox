package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("Exibir o snippet com id %d", id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Exibir formulário para criar um novo snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Salvar um novo snippet..."))
}
