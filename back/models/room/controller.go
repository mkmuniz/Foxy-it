package room

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"w2g-personal-project/utils/errors"

	"github.com/go-chi/chi/v5"
)

func GetAllRoomsController(w http.ResponseWriter, r *http.Request) {
	rooms, err := GetAllRoomsService()

	errors.HandleGetRoom(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)

}

func GetRoomController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	room, err := GetRoomService(int64(id))

	errors.HandleGetRoom(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func CreateRoomController(w http.ResponseWriter, r *http.Request) {
	var model Room

	err := json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	id, err := CreateRoomService(model)

	errors.HandleInsertRoom(w, r, err, id)
}

func PatchRoomController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	errors.HandleConvertString(w, r, err)

	var model Room

	err = json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	rows, err := PatchRoomService(int64(id), model)

	errors.HandleUpdateRoom(w, r, err, rows)
}

func DeleteRoomController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	errors.HandleConvertId(w, r, err)

	errors.HandleDecodeJson(w, r, err)

	rows, err := DeleteRoomService(int64(id))
	errors.HandleDeleteRoom(w, r, err, rows)

}
