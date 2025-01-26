package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// ROUTES
	r.HandleFunc("/", homeHandler).Methods("GET")

	// SERVER
	log.Println("Server is running on http://localhost:8080")
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Fatalf("Could not start server: %s\n", err)
		}
}

// HANDLERS
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(home()).ServeHTTP(w, r)
}
