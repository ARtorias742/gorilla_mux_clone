package main

import (
	"fmt"
	"net/http"

	"github.com/ARtorias742/myrouter/router"
)

func main() {
	r := router.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Data submitted!")
	})

	http.ListenAndServe(":8080", r)
}
