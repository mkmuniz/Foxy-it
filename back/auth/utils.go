package auth

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var resp map[string]any

func Authenticate(w http.ResponseWriter, r *http.Request, password bool, row UserLogin) {
	if password {
		result, err := GenerateToken(row)
		if err != nil {
			resp = map[string]any{
				"status":  500,
				"message": "Error on generate token",
				"data":    err,
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}
		resp = map[string]any{
			"status":  200,
			"message": "Authenticated!",
			"data":    result,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"status":  401,
			"message": "Unauthorized!",
			"data":    nil,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
