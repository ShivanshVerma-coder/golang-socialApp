package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {

	//home Route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to app"))
	}).Methods("GET")

	InitializeBlogRoutes(r.PathPrefix("/blogs").Subrouter())
	InitializeUserRoutes(r.PathPrefix("/users").Subrouter())
}
