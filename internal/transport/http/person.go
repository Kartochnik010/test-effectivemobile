package transport

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) InsertPerson(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err(err).Msg("failed to io.ReadAll(r.Body) at handler.InsertPerson")
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var p models.Person
	if err := json.Unmarshal(content, &p); err != nil {
		h.logger.Err(err).Msg("Error unmarshalling content from " + r.URL.String())
		w.WriteHeader(400)
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	h.service.AddData(&p)
	p, err = h.service.Person.InsertPerson(p)
	if err != nil {
		h.logger.Err(err).Msg("failed to h.service.Person.InsertPerson(p)")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func (h *Handler) FindPersonById(w http.ResponseWriter, r *http.Request) {
	idVar := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idVar)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		h.logger.Err(err).Msg("Invalid ID")
		return
	}

	person, err := h.service.Person.FindPersonById(id)
	if err != nil {
		http.Error(w, "Error fetching person", http.StatusInternalServerError)
		h.logger.Err(err).Msg("Error fetching person")
		return
	}

	jsonPerson, err := json.Marshal(person)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		h.logger.Err(err).Msg("Error encoding JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPerson)

}
func (h *Handler) DeletePersonById(w http.ResponseWriter, r *http.Request) {
	idVar := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idVar)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		h.logger.Err(err).Msg("Invalid ID")
		return
	}

	err = h.service.DeletePersonById(id)
	if err != nil {
		http.Error(w, "Error deleting person", http.StatusInternalServerError)
		h.logger.Err(err).Msg("Error deleting person")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdatePersonById(w http.ResponseWriter, r *http.Request) {
	idVar := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idVar)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		h.logger.Err(err).Msg("Invalid ID")
		return
	}

	var person models.Person
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		h.logger.Err(err).Msg("Error decoding JSON")
		return
	}

	updatedPerson, err := h.service.UpdatePersonById(id, person)
	if err != nil {
		http.Error(w, "Error updating person", http.StatusInternalServerError)
		h.logger.Err(err).Msg("Error updating person")
		return
	}

	jsonUpdatedPerson, err := json.Marshal(updatedPerson)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		h.logger.Err(err).Msg("Error encoding JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUpdatedPerson)
}

// returns ([]models.Person, models.Metadata, error) via http.ResponseWriter
func (h *Handler) SearchPerson(w http.ResponseWriter, r *http.Request) {

}
