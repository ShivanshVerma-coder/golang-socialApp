package routes

import (
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/controllers"
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitializeUserRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.CreateUser).Methods("POST")
	r.Handle("", middleware.IsAuthenticated(controllers.GetAllUsers)).Methods("GET")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.GetUser)).Methods("GET")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.UpdateUser)).Methods("PUT")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.DeleteUser)).Methods("DELETE")
}
