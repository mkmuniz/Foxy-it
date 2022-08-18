package room

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/room", http.HandlerFunc(GetAllRoomsController))
	r.Post("/room", http.HandlerFunc(CreateRoomController))
}
