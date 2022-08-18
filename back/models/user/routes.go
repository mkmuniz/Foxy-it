package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/user", http.HandlerFunc(GetAllUsersController))
	r.Get("/user/{id}", http.HandlerFunc(GetUserController))
	r.Post("/user", http.HandlerFunc(CreateUserController))
	r.Patch("/user/{id}", http.HandlerFunc(PatchUserController))
	r.Delete("/user/{id}", http.HandlerFunc(DeleteUserController))
}
