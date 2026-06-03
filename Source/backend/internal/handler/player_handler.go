package handler

import (
	"encoding/json"
	"net/http"
	"sports-backend/Source/backend/internal/middleware"
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/service"
	"strconv"
)

type PlayerHandler struct {
	service *service.PlayerService
}

func NewPlayerHandler(s *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: s}
}

func (h *PlayerHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /players", h.Create)
	mux.HandleFunc("PUT /players/team", h.UpdateTeam)
	mux.HandleFunc("DELETE /players", h.Delete)
	mux.HandleFunc("GET /players/profiles", h.GetProfiles)
}

func (h *PlayerHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var p models.Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreatePlayer(&p, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *PlayerHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var data struct {
		ID     int `json:"id"`
		TeamID int `json:"team_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdatePlayerTeam(data.ID, data.TeamID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *PlayerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.service.DeletePlayer(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *PlayerHandler) GetProfiles(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	profiles, err := h.service.GetPlayerProfiles(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}
