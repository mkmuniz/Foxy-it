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
	users, err := GetAllUsersService()

	errors.HandleGetUser(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	user, err := GetUserService(int64(id))

	errors.HandleGetUser(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var model User

	err := json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	id, err := CreateUserService(model)

	errors.HandleInsertUser(w, r, err, id)
}

func PatchUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	errors.HandleConvertId(w, r, err)

	var model User

	err = json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	rows, err := PatchUserService(int64(id), model)

	errors.HandleUpdateUser(w, r, err)

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

	errors.HandleConvertId(w, r, err)

	rows, err := DeleteUserService(int64(id))

	errors.HandleDeleteUser(w, r, err)

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
