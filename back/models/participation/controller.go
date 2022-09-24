package participation

import (
	"encoding/json"
	"net/http"
	"strconv"
	"w2g-personal-project/utils/errors"

	"github.com/go-chi/chi/v5"
)

func GetAllParticipationController(w http.ResponseWriter, r *http.Request) {
	participations, err := GetAllParticipationsService()

	errors.HandleGetParticipation(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participations)
}

func GetParticipationController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	errors.HandleConvertId(w, r, err)

	participation, err := GetParticipationService(int64(id))

	errors.HandleGetParticipation(w, r, err)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participation)
}

func CreateParticipationController(w http.ResponseWriter, r *http.Request) {
	var model Participation

	err := json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	id, err := CreateParticipationService(model)

	errors.HandleInsertParticipation(w, r, err, id)
}

func PatchParticipationController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	errors.HandleConvertString(w, r, err)

	var model Participation

	err = json.NewDecoder(r.Body).Decode(&model)

	errors.HandleDecodeJson(w, r, err)

	rows, err := PatchParticipationService(int64(id), model)

	errors.HandleUpdateParticipation(w, r, err, rows)
}

func DeleteParticipationController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	errors.HandleConvertId(w, r, err)

	errors.HandleDecodeJson(w, r, err)

	rows, err := DeleteParticipationService(int64(id))

	errors.HandleDeleteParticipation(w, r, err, rows)
}
