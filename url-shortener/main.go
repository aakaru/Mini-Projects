package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL struct represents a shortened URL entry
type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// In-memory database for storing URLs
var urlDB = make(map[string]URL)

// generateShortURL creates a short hash from the original URL
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL + time.Now().String())) // Add timestamp for uniqueness
	hash := hex.EncodeToString(hasher.Sum(nil))[:8]        // Use first 8 chars
	return hash
}

// createURL stores the original URL and returns a shortened URL
func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	urlDB[shortURL] = URL{
		Id:           shortURL,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// getURL retrieves the original URL from the database
func getURL(id string) (URL, error) {
	url, exists := urlDB[id]
	if !exists {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

// handler serves a simple welcome message
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the URL Shortener Service!")
}

// shortURLHandler handles URL shortening requests
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		URL string `json:"url"`
	}

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Ensure the URL is not empty
	if requestData.URL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate the short URL
	shortURL := createURL(requestData.URL)

	// Send response
	responseData := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: "http://localhost:3000/redirect/" + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// redirectURLHandler redirects short URLs to their original URLs
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// Define routes
	http.HandleFunc("/", handler)
	http.HandleFunc("/shortener", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server
	port := ":3000"
	fmt.Println("Starting server on PORT", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
