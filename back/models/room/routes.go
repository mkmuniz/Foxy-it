package room

import (
	"github.com/go-chi/chi/v5"
)

func RoomRoutes() {
	r := chi.NewRouter()

	r.Get("/", GetAllRoomsController())
	r.Post("/", CreateRoomController())
	r.Patch("/", PatchRoomController())
	r.Delete("/", DeleteRoomController())
}
