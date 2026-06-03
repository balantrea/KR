package handler

import (
	"encoding/json"
	"net/http"
	"sports-backend/Source/backend/internal/middleware"
	"sports-backend/Source/backend/internal/models"
	"sports-backend/Source/backend/internal/service"
)

type MatchHandler struct {
	service *service.MatchService
}

func NewMatchHandler(s *service.MatchService) *MatchHandler {
	return &MatchHandler{service: s}
}

func (h *MatchHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /matches", h.Create)
}

func (h *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	var m models.Match
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateMatch(&m, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
