package main

import (
	"contact-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Register routes from the routes package
	routes.RegisterRoutes(router)

	// Set up CORS middleware to allow requests from your Vue frontend (localhost:5173)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Replace with your frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the server with CORS handling
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
