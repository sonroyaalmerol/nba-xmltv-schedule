package main

import (
	"encoding/xml"
	"fmt"
	"nba-xmltv-schedule/nba"
	"nba-xmltv-schedule/xmltv"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ServeNBASchedule serves the NBA schedule in XMLTV format
func ServeNBASchedule(w http.ResponseWriter, r *http.Request) {
	nbaSchedule, err := nba.FetchNBASchedule()
	if err != nil {
		http.Error(w, "Error fetching NBA schedule", http.StatusInternalServerError)
		return
	}

	iconBaseURL := ""
	if r.TLS == nil {
		iconBaseURL = fmt.Sprintf("http://%s/logos", r.Host)
	} else {
		iconBaseURL = fmt.Sprintf("https://%s/logos", r.Host)
	}

	xmlSchedule := xmltv.ConvertToXMLTV(nbaSchedule, iconBaseURL)

	// Convert schedule to XML
	xmlData, err := xml.MarshalIndent(xmlSchedule, "", "  ")
	if err != nil {
		http.Error(w, "Error creating XML", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/xml")

	// Add XML declaration and DOCTYPE
	xmlHeader := `<?xml version="1.0" encoding="ISO-8859-1"?>
		<!DOCTYPE tv SYSTEM "xmltv.dtd">` + "\n"

	// Write the XML declaration and DOCTYPE followed by the actual XML data to the response writer
	w.Write([]byte(xmlHeader))

	// Write XML data to response
	w.Write(xmlData)
}

func ServeLogo(w http.ResponseWriter, r *http.Request) {
	// Extract the channel ID from the URL path
	channelID := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/logos/"), ".png")
	if channelID == "" {
		http.NotFound(w, r)
		return
	}

	// Construct the file path based on the channel ID
	filePath := filepath.Join("logos", fmt.Sprintf("%s.png", channelID))

	// Read the image file
	image, err := os.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set the Content-Type header to image/png
	w.Header().Set("Content-Type", "image/png")

	// Write the image data to the response writer
	w.Write(image)
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/nba_schedule", ServeNBASchedule)
	http.HandleFunc("/logos/", ServeLogo)

	// Start the server
	port := 8080
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost:%d/nba_schedule\n", port)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
