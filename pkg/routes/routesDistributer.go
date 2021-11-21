package routes

import (
	"net/http"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/controllers"
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {

	//home Route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to app"))
	}).Methods("GET")

	//login
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	r.Handle("/profile", middleware.IsAuthenticated(controllers.GetProfile)).Methods("GET")

	//transfer to controllers by distributing
	InitializeBlogRoutes(r.PathPrefix("/blogs").Subrouter())
	InitializeUserRoutes(r.PathPrefix("/users").Subrouter())
}
