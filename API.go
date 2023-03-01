package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func TimeHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Method is not supported, use GET request", http.StatusNotFound)
		return
	} else if req.URL.Path != "/time" {
		http.Error(writer, "Invalid API call", http.StatusNotFound)
		return
	}

	TimeResponse(writer)
}

func TimeResponse(resp http.ResponseWriter) {
	time := time.Now().Format(time.RFC3339)
	body := map[string]string{"time": time}
	resp.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(resp).Encode(body)
	if err != nil {
		log.Fatalf("Api error. Err: %s", err)
	}
}
