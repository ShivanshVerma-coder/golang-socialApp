package routes

import (
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/controllers"
	"github.com/gorilla/mux"
)

func InitializeBlogRoutes(r *mux.Router) {
	r.HandleFunc("", controllers.GetAllBlogs).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetBlog).Methods("GET")
	r.HandleFunc("", controllers.CreateBlog).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateBlog).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteBlog).Methods("DELETE")
}
