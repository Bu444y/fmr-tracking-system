package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
	"github.com/Bu444y/fmr-tracking-system/internal/repository"
	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
	repo *repository.NoteRepository
}

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		repo: repository.NewNoteRepository(),
	}
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.repo.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) List(w http.ResponseWriter, r *http.Request) {
	fmrIDStr := chi.URLParam(r, "fmrId")
	fmrID, err := strconv.ParseInt(fmrIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid FMR ID", http.StatusBadRequest)
		return
	}

	response, err := h.repo.ListByFMRID(r.Context(), fmrID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
