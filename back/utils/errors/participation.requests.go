package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGetParticipation(w http.ResponseWriter, r *http.Request, err error) (res error) {
	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error, participation not found!",
			"data":    err,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		return err
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Participation found!",
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		return nil
	}
}

func HandleInsertParticipation(w http.ResponseWriter, r *http.Request, err error, id int64) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": fmt.Sprintf("Error on insert participation: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"status":  201,
			"message": fmt.Sprintf("Participation inserted successfully"),
			"id":      id,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleUpdateParticipation(w http.ResponseWriter, r *http.Request, err error, rows int64) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on update participation: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Participation updated successfully"),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	if rows > 1 {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Many participations was updated: %v", rows),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleDeleteParticipation(w http.ResponseWriter, r *http.Request, err error, rows int64) {
	if err != nil {
		resp = map[string]any{
			"error":   500,
			"message": fmt.Sprintf("Error on delete Participation: %v", err),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Participation deleted successfully"),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	if rows > 1 {
		resp = map[string]any{
			"error":   200,
			"message": fmt.Sprintf("Many participations was deleted: %v", rows),
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
