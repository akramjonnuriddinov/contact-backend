package routes

import (
	"contact-backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/contacts", controllers.GetAllContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", controllers.GetContactByID).Methods("GET")
	router.HandleFunc("/contacts", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", controllers.DeleteContact).Methods("DELETE")
}
