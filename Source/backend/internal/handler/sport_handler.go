package handler

import (
	"encoding/json"
	"net/http"
	"sports-backend/Source/backend/internal/middleware"
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/service"
	"strconv"
)

type SportHandler struct {
	service *service.SportService
}

func NewSportHandler(s *service.SportService) *SportHandler {
	return &SportHandler{service: s}
}

func (h *SportHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /sports", h.Create)
	mux.HandleFunc("GET /sports", h.GetAll)
	mux.HandleFunc("PUT /sports", h.Update)
	mux.HandleFunc("DELETE /sports", h.Delete)
	mux.HandleFunc("GET /sports/count", h.GetTeamCount)
}

func (h *SportHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var s models.Sport
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateSport(&s, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func (h *SportHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	sports, err := h.service.GetAllSports(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sports)
}

func (h *SportHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var s models.Sport
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateSport(&s, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *SportHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteSport(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SportHandler) GetTeamCount(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	count, err := h.service.GetSportTeamCount(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"team_count": count})
}
