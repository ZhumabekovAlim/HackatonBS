package handlers

import (
	"BS_Hackathon/internal/models"
	"BS_Hackathon/internal/services"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	Service *services.BookService
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // default to page 1 if no valid page is provided
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // default to 10 books per page if no valid limit is provided
	}

	books, err := h.Service.GetAllBooks(r.Context(), page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.Service.GetBookByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdBook, err := h.Service.CreateBook(r.Context(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	book.ID = id

	updatedBook, err := h.Service.UpdateBook(r.Context(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(":id")
	if idStr == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteBook(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookHandler) PhotoDetect(w http.ResponseWriter, r *http.Request) {
	// Parse form data to get the uploaded file
	err := r.ParseMultipartForm(10 << 20) // limit max file size to 10MB
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		log.Printf("Error retrieving file: %v", err)
		return
	}
	defer file.Close()

	// Use the full Cloudinary URL from your environment
	cld, err := cloudinary.NewFromURL("cloudinary://697114559996688:Iu6d48qksAJc95d1aBAXv0-jqg8@dv9gjjh73")
	if err != nil {
		http.Error(w, "Error initializing Cloudinary client", http.StatusInternalServerError)
		log.Printf("Error initializing Cloudinary: %v", err)
		return
	}

	// Upload file to Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: strings.TrimSuffix(handler.Filename, ".png"), // Use filename as public ID
		Folder:   "uploads",                                    // Folder in Cloudinary
	})
	if err != nil {
		http.Error(w, "Error uploading image to Cloudinary", http.StatusInternalServerError)
		log.Printf("Error uploading image: %v", err)
		return
	}

	// Get the secure URL of the uploaded image
	imageURL := uploadResult.SecureURL
	log.Printf("Image uploaded successfully to Cloudinary: %s", imageURL)

	// Call AI API to analyze the photo using the Cloudinary image URL
	recognizedText, err := h.AnalyzePhoto(imageURL)
	if err != nil {
		http.Error(w, "Error analyzing photo", http.StatusInternalServerError)
		log.Printf("Error analyzing photo: %v", err)
		return
	}

	// Search for books based on the recognized text from the AI API
	books, err := h.Service.FindBooksByTitle(r.Context(), recognizedText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error finding books: %v", err)
		return
	}

	// Return the found books in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// AnalyzePhoto sends the image URL to the OpenAI API for analysis and returns the recognized text
func (h *BookHandler) AnalyzePhoto(imageURL string) (string, error) {
	// OpenAI API URL and Key from environment variables
	url := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"

	// Create the payload for the OpenAI API call
	payload := map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": "What's in this image? Write only one main word",
					},
					{
						"type":      "image_url",
						"image_url": map[string]string{"url": imageURL},
					},
				},
			},
		},
		"max_tokens": 300,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Printf("OpenAI response body: %s", string(body)) // Log the full response for debugging

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error from OpenAI: status code %d, response: %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("OpenAI API returned status code %d", resp.StatusCode)
	}

	// Parse the response
	var openAIResponse OpenAIResponse
	if err := json.Unmarshal(body, &openAIResponse); err != nil {
		return "", err
	}

	// Return the recognized text
	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}

func (h *BookHandler) GetAllNewBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllNewBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
