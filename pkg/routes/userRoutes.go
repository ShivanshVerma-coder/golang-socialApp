package routes

import (
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/controllers"
	"github.com/gorilla/mux"
)

func InitializeUserRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
}
