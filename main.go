package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const hostAdress = ":3000"
const clientAdress = "http://127.0.0.1:5173"

func main() {

	r := chi.NewRouter()

	r.Use(Cors)
	r.Use(middleware.Logger)

	r.Mount("/api", emailsResource{}.Routes())

	log.Printf("Listening: " + hostAdress)
	log.Fatal(http.ListenAndServe(hostAdress, r))
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Allow-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Allow-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
