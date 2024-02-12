package main

import (
	"log"
	"net/http"
)

// var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando páginas de usuários!"))
}

func main() {

	// templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
