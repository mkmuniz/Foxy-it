package room

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAllRoomsController(w http.ResponseWriter, r *http.Request) {
	rooms, err := GetAllRoomsService()
	if err != nil {
		log.Printf("Error on get all rooms: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)

}

func CreateRoomController(w http.ResponseWriter, r *http.Request) {
	var model Room

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := CreateRoomService(model)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   500,
			"Message": fmt.Sprintf("Error on insert user: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   201,
			"Message": fmt.Sprintf("User inserted successfully, ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return
}
