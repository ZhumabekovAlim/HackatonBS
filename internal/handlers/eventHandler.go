package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type EventHandler struct {
	Service *services.EventService
}

func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.Service.GetAllEvents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := h.Service.GetEventByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdEvent, err := h.Service.CreateEvent(r.Context(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEvent)
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	var event models.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	event.ID = id

	updatedEvent, err := h.Service.UpdateEvent(r.Context(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing event ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteEvent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
