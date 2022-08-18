package server

import (
	"log"
	"net/http"
	"w2g-personal-project/auth"
	"w2g-personal-project/configs"
	"w2g-personal-project/models/room"
	"w2g-personal-project/models/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Back is running"))
	}))
	r.Post("/login", auth.LoginController)
	r.Group(room.Routes)
	r.Group(user.Routes)

	log.Printf("Server started on port%v", configs.GetAPIConfig())
	log.Fatal(http.ListenAndServe(configs.GetAPIConfig(), r))
}
