package participation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/participation", http.HandlerFunc(GetAllParticipationController))
	r.Get("/participation/{id}", http.HandlerFunc(GetParticipationController))
	r.Post("/participation", http.HandlerFunc(CreateParticipationController))
	r.Put("/participation/{id}", http.HandlerFunc(PatchParticipationController))
	r.Delete("/participation/{id}", http.HandlerFunc(DeleteParticipationController))
}
