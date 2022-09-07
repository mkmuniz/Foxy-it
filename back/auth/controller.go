package auth

import (
	"encoding/json"
	"net/http"
	"w2g-personal-project/utils/errors"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	var user UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)

	errors.HandleDBConnection(w, r, err)

	row, err := LoginService(user)

	errors.HandleGetUser(w, r, err)

	isPasswordValid := CheckPasswordHash(user.Password, row.Password)

	VerifyPassword(w, r, isPasswordValid, row)
}
