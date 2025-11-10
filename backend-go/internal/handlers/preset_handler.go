package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
	"github.com/Bu444y/fmr-tracking-system/internal/repository"
	"github.com/go-chi/chi/v5"
)

type PresetHandler struct {
	repo *repository.PresetRepository
}

func NewPresetHandler() *PresetHandler {
	return &PresetHandler{
		repo: repository.NewPresetRepository(),
	}
}

// Filter Presets
func (h *PresetHandler) CreateFilterPreset(w http.ResponseWriter, r *http.Request) {
	var req models.CreatePresetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	preset, err := h.repo.CreateFilterPreset(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preset)
}

func (h *PresetHandler) ListFilterPresets(w http.ResponseWriter, r *http.Request) {
	response, err := h.repo.ListFilterPresets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PresetHandler) DeleteFilterPreset(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteFilterPreset(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Column Presets
func (h *PresetHandler) CreateColumnPreset(w http.ResponseWriter, r *http.Request) {
	var req models.CreateColumnPresetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	preset, err := h.repo.CreateColumnPreset(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preset)
}

func (h *PresetHandler) ListColumnPresets(w http.ResponseWriter, r *http.Request) {
	response, err := h.repo.ListColumnPresets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PresetHandler) DeleteColumnPreset(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteColumnPreset(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
