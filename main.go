package main

import (
	"fmt"
	"net/http"

	"github.com/ARtorias742/matcher"
	"github.com/ARtorias742/middleware"
	"github.com/ARtorias742/router"
)

func main() {
	r := router.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	}).Name("home").Host("example.com")

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := matcher.GetParams(r.Context())
		fmt.Fprintf(w, "User ID: %s", params["id"])
	}).Use(middleware.Logging)

	r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Submitted")
	})

	http.ListenAndServe(":8080", r)
}
