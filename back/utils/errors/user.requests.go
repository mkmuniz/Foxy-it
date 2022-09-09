package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleConvertId(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on convert ID!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleDecodeJson(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on deconding json!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleGetUser(w http.ResponseWriter, r *http.Request, err error) (res error) {
	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error, user not found!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		return err
	}

	return
}

func HandleInsertUser(w http.ResponseWriter, r *http.Request, err error, id int64) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": fmt.Sprintf("Error on insert user: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"status":  201,
			"message": fmt.Sprintf("User inserted successfully, ID: %v", id),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on update user: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on delete user: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
