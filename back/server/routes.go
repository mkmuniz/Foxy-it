package server

import (
	"log"
	"net/http"
	"w2g-personal-project/models/room"
	"w2g-personal-project/models/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(room.Routes)
	r.Group(user.Routes)

	log.Printf("Server started on port %v", ":8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
