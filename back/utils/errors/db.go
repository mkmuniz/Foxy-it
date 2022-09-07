package errors

import (
	"encoding/json"
	"net/http"
)

var resp map[string]any

func HandleDBConnection(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on connect to database",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
