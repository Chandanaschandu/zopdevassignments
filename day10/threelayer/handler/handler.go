package handler

import (
	"encoding/json"
	"github.com/Chandanaschandu/threelayer/models"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type handler struct {
	service UserServiceInterface
}

func NewUserHandler(userService UserServiceInterface) *handler {
	return &handler{service: userService}
}

func (h *handler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["user_name"]
	user, err := h.service.GetUserByName(userName)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *handler) AddUsers(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var NewUser models.Users
	err = json.Unmarshal(body, &NewUser)

	if err != nil {
		http.Error(w, "Error unmarshalling", http.StatusBadRequest)
	}
	err = NewUser.Validate()

	if err != nil {
		http.Error(w, "Error in validation email or phone number", http.StatusBadRequest)
		return
	}

	err = h.service.AddUser(&NewUser)

	if err != nil {
		http.Error(w, "Error adding user", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(&NewUser)

	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Username := vars["name"]
	err := h.service.DeleteUser(Username)

	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (h *handler) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var user models.Users

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if len(body) > 0 {
		err := json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "Invalid user data", http.StatusBadRequest)
			return
		}
	}

	err = user.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateUserEmail(name, user.Email)
	if err != nil {
		http.Error(w, "Error updating the email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("data updated"))

}
