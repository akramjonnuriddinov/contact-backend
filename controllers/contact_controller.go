package controllers

import (
	"contact-backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var contacts []models.Contact

// GetAllContacts - Fetches all contacts
func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

// GetContactByID - Fetches a single contact by its ID
func GetContactByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, contact := range contacts {
		if contact.ID == id {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}

// CreateContact - Adds a new contact
func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	json.NewDecoder(r.Body).Decode(&contact)

	contact.ID = len(contacts) + 1
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// UpdateContact - Updates an existing contact by ID
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, contact := range contacts {
		if contact.ID == id {
			var updatedContact models.Contact
			json.NewDecoder(r.Body).Decode(&updatedContact)
			contacts[i] = updatedContact
			contacts[i].ID = id
			json.NewEncoder(w).Encode(contacts[i])
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}

// DeleteContact - Deletes a contact by ID
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			json.NewEncoder(w).Encode(contacts)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}
