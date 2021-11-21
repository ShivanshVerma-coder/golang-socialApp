package routes

import (
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/controllers"
	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/middleware"
	"github.com/gorilla/mux"
)

func InitializeBlogRoutes(r *mux.Router) {
	r.Handle("", middleware.IsAuthenticated(controllers.GetAllBlogs)).Methods("GET")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.GetBlog)).Methods("GET")
	r.Handle("", middleware.IsAuthenticated(controllers.CreateBlog)).Methods("POST")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.UpdateBlog)).Methods("PUT")
	r.Handle("/{id}", middleware.IsAuthenticated(controllers.DeleteBlog)).Methods("DELETE")
}
