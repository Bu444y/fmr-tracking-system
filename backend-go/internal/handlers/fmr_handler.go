package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bu444y/fmr-tracking-system/internal/models"
	"github.com/Bu444y/fmr-tracking-system/internal/repository"
	"github.com/go-chi/chi/v5"
)

type FMRHandler struct {
	repo *repository.FMRRepository
}

func NewFMRHandler() *FMRHandler {
	return &FMRHandler{
		repo: repository.NewFMRRepository(),
	}
}

func (h *FMRHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateFMRRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmr, err := h.repo.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmr)
}

func (h *FMRHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	fmr, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmr)
}

func (h *FMRHandler) List(w http.ResponseWriter, r *http.Request) {
	filters := &models.FMRFilters{}

	// Parse query parameters
	if status := r.URL.Query().Get("status"); status != "" {
		s := models.FMRStatus(status)
		filters.Status = &s
	}
	if search := r.URL.Query().Get("search"); search != "" {
		filters.Search = &search
	}
	if sortBy := r.URL.Query().Get("sortBy"); sortBy != "" {
		filters.SortBy = &sortBy
	}
	if sortOrder := r.URL.Query().Get("sortOrder"); sortOrder != "" {
		filters.SortOrder = &sortOrder
	}
	if limit := r.URL.Query().Get("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filters.Limit = l
		}
	}
	if offset := r.URL.Query().Get("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filters.Offset = o
		}
	}
	if tierLevel := r.URL.Query().Get("tierLevel"); tierLevel != "" {
		tl := models.TierLevel(tierLevel)
		filters.TierLevel = &tl
	}

	response, err := h.repo.List(r.Context(), filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *FMRHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateFMRRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ID = id

	fmr, err := h.repo.Update(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmr)
}
