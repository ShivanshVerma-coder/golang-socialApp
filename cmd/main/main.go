package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Main ran")

	r := mux.NewRouter() //router

	//Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true, // set this to true if your api supports cookies
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(r)

	routes.InitializeRoutes(r)                       // Initialize all routes
	log.Fatal(http.ListenAndServe(":8080", handler)) // start server
}
