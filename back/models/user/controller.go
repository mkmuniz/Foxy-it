package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"w2g-personal-project/utils/errors"

	"github.com/go-chi/chi/v5"
)

func GetAllUsersController(w http.ResponseWriter, r *http.Request) {

	rooms, err := GetAllUsersService()
	if err != nil {
		log.Printf("Error on get all users: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)

}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := GetUserService(int64(id))

	errors.HandleGetUser(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var model User

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := CreateUserService(model)

	errors.HandleInsertUser(w, r, err, id)
}

func PatchUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var model User

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := PatchUserService(int64(id), model)
	if err != nil {
		log.Printf("Error on update user: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Many users was update: %v", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   200,
		"Message": fmt.Sprintf("User updated successfully, ID: %v", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := DeleteUserService(int64(id))
	if err != nil {
		log.Printf("Error on delete user: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Many users was deleted: %v", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   200,
		"Message": fmt.Sprintf("User deleted successfully, ID: %v", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
