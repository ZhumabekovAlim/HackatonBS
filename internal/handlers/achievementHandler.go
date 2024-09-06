package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type AchievementHandler struct {
	Service *services.AchievementService
}

func (h *AchievementHandler) GetAllAchievements(w http.ResponseWriter, r *http.Request) {
	achievements, err := h.Service.GetAllAchievements(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}

func (h *AchievementHandler) GetAchievementByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid achievement ID", http.StatusBadRequest)
		return
	}

	achievement, err := h.Service.GetAchievementByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievement)
}

func (h *AchievementHandler) CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	err := json.NewDecoder(r.Body).Decode(&achievement)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdAchievement, err := h.Service.CreateAchievement(r.Context(), achievement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAchievement)
}

func (h *AchievementHandler) UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid achievement ID", http.StatusBadRequest)
		return
	}

	var achievement models.Achievement
	err = json.NewDecoder(r.Body).Decode(&achievement)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	achievement.ID = id

	updatedAchievement, err := h.Service.UpdateAchievement(r.Context(), achievement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAchievement)
}

func (h *AchievementHandler) DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid achievement ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteAchievement(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
