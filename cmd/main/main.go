package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main ran")
	r := mux.NewRouter()

	routes.InitializeRoutes(r)                 // Initialize all routes
	log.Fatal(http.ListenAndServe(":8080", r)) // start server
}
