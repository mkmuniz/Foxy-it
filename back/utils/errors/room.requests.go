package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGetRoom(w http.ResponseWriter, r *http.Request, err error) (res error) {
	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error, Room not found!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		return err
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Room found!",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		return nil
	}

	return
}

func HandleInsertRoom(w http.ResponseWriter, r *http.Request, err error, id int64) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": fmt.Sprintf("Error on insert Room: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"status":  201,
			"message": fmt.Sprintf("Room inserted successfully, ID: %v", id),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleConvertString(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error converting ID: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleUpdateRoom(w http.ResponseWriter, r *http.Request, err error, rows int64) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on update Room: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Room updated successfully"),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	if rows > 1 {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Many rooms was updated: %v", rows),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleDeleteRoom(w http.ResponseWriter, r *http.Request, err error, rows int64) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on delete Room: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Room deleted successfully"),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	if rows > 1 {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Many rooms was deleted: %v", rows),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
