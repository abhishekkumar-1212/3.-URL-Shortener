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

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// In-memory data structure
var urlDB = make(map[string]URL)

// Using MD5 hashing algorithm to generate a short URL
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL)) // Convert URL into byte slice
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:10]
}


func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL // Use the short URL as the ID for simplicity
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}


func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok { // If not exist inside the map
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}


func rootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello bros Cp")
}


func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	type requestData struct {
		URL string `json:"url"`
	}

	var data requestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)
	responseData := map[string]string{"short_url": shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}


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
	
	fmt.Println("Starting URL Shortener...")

	// Register the handler functions
	http.HandleFunc("/", rootPageURL)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
