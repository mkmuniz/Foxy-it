package room

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/room", http.HandlerFunc(GetAllRoomsController))
	r.Get("/room/{id}", http.HandlerFunc(GetRoomController))
	r.Post("/room", http.HandlerFunc(CreateRoomController))
	r.Patch("/room/{id}", http.HandlerFunc(PatchRoomController))
	r.Delete("/room/{id}", http.HandlerFunc(DeleteRoomController))
}
