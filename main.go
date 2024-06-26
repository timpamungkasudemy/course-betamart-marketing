package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

const (
	HttpResponseHeaderContentType = "Content-Type"
	MimeApplicationJSON           = "application/json"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthData := map[string]string{
		"status": "up",
	}
	jsonResponse, err := json.Marshal(healthData)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set(HttpResponseHeaderContentType, MimeApplicationJSON)
	w.Write(jsonResponse)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04:05")

	timeData := map[string]string{
		"time": formattedTime,
	}

	jsonResponse, err := json.Marshal(timeData)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set(HttpResponseHeaderContentType, MimeApplicationJSON)
	w.Write(jsonResponse)
}

var buzzwords = []string{
	"Synergy", "Leverage", "Pivoting", "Holistic Approach", "Paradigm Shift",
	"Disruptive Innovation", "Thought Leadership", "Agile", "Big Data", "Blockchain",
}

func buzzwordHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(uint64(time.Now().UnixNano()))

	selectedIndices := rand.Perm(len(buzzwords))[:3]
	selectedBuzzwords := make([]string, 0, 3)
	for _, index := range selectedIndices {
		selectedBuzzwords = append(selectedBuzzwords, buzzwords[index])
	}

	response, err := json.Marshal(selectedBuzzwords)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set(HttpResponseHeaderContentType, MimeApplicationJSON)
	w.Write(response)
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/api/time", timeHandler)
	http.HandleFunc("/api/buzzwords", buzzwordHandler)

	fmt.Println("Server starting on port 8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
