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

	teste := errors.HandleGetUser(w, r, err)

	if teste != nil {
		return
	}

	isPasswordValid := CheckPasswordHash(user.Password, row.Password)

	Authenticate(w, r, isPasswordValid, row)
}
