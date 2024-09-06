package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserAchievementHandler struct {
	Service *services.UserAchievementService
}

func (h *UserAchievementHandler) GetAllUserAchievements(w http.ResponseWriter, r *http.Request) {
	userAchievements, err := h.Service.GetAllUserAchievements(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userAchievements)
}

func (h *UserAchievementHandler) GetUserAchievementByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_achievement ID", http.StatusBadRequest)
		return
	}

	userAchievement, err := h.Service.GetUserAchievementByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userAchievement)
}

func (h *UserAchievementHandler) CreateUserAchievement(w http.ResponseWriter, r *http.Request) {
	var userAchievement models.UserAchievement
	err := json.NewDecoder(r.Body).Decode(&userAchievement)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUserAchievement, err := h.Service.CreateUserAchievement(r.Context(), userAchievement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUserAchievement)
}

func (h *UserAchievementHandler) UpdateUserAchievement(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_achievement ID", http.StatusBadRequest)
		return
	}

	var userAchievement models.UserAchievement
	err = json.NewDecoder(r.Body).Decode(&userAchievement)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userAchievement.ID = id

	updatedUserAchievement, err := h.Service.UpdateUserAchievement(r.Context(), userAchievement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUserAchievement)
}

func (h *UserAchievementHandler) DeleteUserAchievement(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_achievement ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_achievement ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUserAchievement(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
