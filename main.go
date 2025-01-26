package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// ROUTES
	r.HandleFunc("/", homeHandler).Methods("GET")

	// SERVER
	http.ListenAndServe(":8080", r)
}

// HANDLERS
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(home()).ServeHTTP(w, r)
}