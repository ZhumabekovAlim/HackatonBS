package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserEventHandler struct {
	Service *services.UserEventService
}

func (h *UserEventHandler) GetAllUserEvents(w http.ResponseWriter, r *http.Request) {
	userEvents, err := h.Service.GetAllUserEvents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userEvents)
}

func (h *UserEventHandler) GetUserEventByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_event ID", http.StatusBadRequest)
		return
	}

	userEvent, err := h.Service.GetUserEventByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userEvent)
}

func (h *UserEventHandler) CreateUserEvent(w http.ResponseWriter, r *http.Request) {
	var userEvent models.UserEvent
	err := json.NewDecoder(r.Body).Decode(&userEvent)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUserEvent, err := h.Service.CreateUserEvent(r.Context(), userEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUserEvent)
}

func (h *UserEventHandler) UpdateUserEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_event ID", http.StatusBadRequest)
		return
	}

	var userEvent models.UserEvent
	err = json.NewDecoder(r.Body).Decode(&userEvent)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userEvent.ID = id

	updatedUserEvent, err := h.Service.UpdateUserEvent(r.Context(), userEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUserEvent)
}

func (h *UserEventHandler) DeleteUserEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing user_event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_event ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUserEvent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
