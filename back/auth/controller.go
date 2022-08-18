package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	var user UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	row, err := LoginService(user)

	if err != nil {
		log.Printf("Error on search user: %v", err)
	}

	resp := CheckPasswordHash(user.Password, row.Password)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
