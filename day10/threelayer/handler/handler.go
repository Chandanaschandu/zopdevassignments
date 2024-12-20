package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Chandanaschandu/threelayer/models"
	"github.com/Chandanaschandu/threelayer/service"
)

type Handler struct {
	service *service.UserService
}

func NewUserHandler(userService *service.UserService) Handler {
	return Handler{service: userService}
}

func (h *Handler) GetUserByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	switch r.Method {
	case http.MethodGet:
		user, err := h.service.GetUserByName(name)
		if err != nil {
			http.Error(w, "Error fetching user", http.StatusInternalServerError)
			return
		}

		if user.UserName == "" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)

		if err != nil {
			http.Error(w, "Error marshaling response", http.StatusInternalServerError)
			return

		}

	case http.MethodPut:
		var user models.Users

		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		if len(body) > 0 {
			err := json.Unmarshal(body, &user)
			if err != nil {
				fmt.Println()
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
			fmt.Println()
			switch err {

			}
			http.Error(w, "Error in updating the email", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("data updated"))

	case http.MethodDelete:
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

}

func (h *Handler) AddUsers(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var NewUser models.Users
	err = json.Unmarshal(body, &NewUser)

	if err != nil {
		http.Error(w, "Error in unmarshalling", http.StatusBadRequest)
	}
	err = NewUser.Validate()

	if err != nil {
		http.Error(w, "error in validation email or phone number", http.StatusBadRequest)
		return
	}

	err = h.service.AddUsers(&NewUser)

	if err != nil {
		http.Error(w, "error in adding users", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(&NewUser)

	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
