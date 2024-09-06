package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserBookHandler struct {
	Service *services.UserBookService
}

func (h *UserBookHandler) GetAllUserBooks(w http.ResponseWriter, r *http.Request) {
	userBooks, err := h.Service.GetAllUserBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userBooks)
}

func (h *UserBookHandler) GetUserBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing user book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user book ID", http.StatusBadRequest)
		return
	}

	userBook, err := h.Service.GetUserBookByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userBook)
}

func (h *UserBookHandler) CreateUserBook(w http.ResponseWriter, r *http.Request) {
	var userBook models.UserBook
	err := json.NewDecoder(r.Body).Decode(&userBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUserBook, err := h.Service.CreateUserBook(r.Context(), userBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUserBook)
}

func (h *UserBookHandler) UpdateUserBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing user book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user book ID", http.StatusBadRequest)
		return
	}

	var userBook models.UserBook
	err = json.NewDecoder(r.Body).Decode(&userBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userBook.ID = id

	updatedUserBook, err := h.Service.UpdateUserBook(r.Context(), userBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUserBook)
}

func (h *UserBookHandler) DeleteUserBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing user book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user book ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUserBook(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
