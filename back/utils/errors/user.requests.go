package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error, user not found!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleInsertUser(w http.ResponseWriter, r *http.Request, err error, id int64) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on insert user: %v", err),
		}
	} else {
		resp = map[string]any{
			"error":   201,
			"Message": fmt.Sprintf("User inserted successfully, ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
