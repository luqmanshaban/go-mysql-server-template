package routes

import (
	"net/http"
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	// Create a subrouter for routes that need the middleware
	sr := r.PathPrefix("").Subrouter()

	sr.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/users", controllers.SignUp).Methods("POST").Host("localhost:3000")
	r.HandleFunc("/", controllers.Hello)

	http.Handle("/", r)
}