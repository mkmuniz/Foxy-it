package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var resp map[string]any

func VerifyPassword(w http.ResponseWriter, r *http.Request, password bool, row UserLogin) {
	log.Print(row)
	if password {
		result, err := GenerateToken(row)
		if err != nil {
			resp = map[string]any{
				"status":  500,
				"message": "Error on generate token",
				"data":    err,
			}
		}

		resp = map[string]any{
			"status":  200,
			"message": "Token generated",
			"data":    result,
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
