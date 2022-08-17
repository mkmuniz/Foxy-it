package room

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RoomRoutes(w http.ResponseWriter, r *http.Request) {
	routess := chi.NewRouter()

	routess.Get("/teste/", GetAllRoomsController)
	routess.Post("/teste/", CreateRoomController)
}
