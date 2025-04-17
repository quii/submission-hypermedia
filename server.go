package main

import (
	"net/http"
)

// SetupServer configures and returns an HTTP server
func SetupServer() *http.Server {
	handler := NewAppHandler()

	// Set up routes
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Page routes
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/pages/", handler.PageHandler)

	// Component routes
	mux.HandleFunc("/components/submission", handler.SubmissionComponentHandler)

	// Action routes
	mux.HandleFunc("/submission/approve", handler.ApproveSubmissionHandler)
	mux.HandleFunc("/submission/reject", handler.RejectSubmissionHandler)
	mux.HandleFunc("/submission/toggle-watchlist", handler.ToggleWatchlistHandler)

	// Create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
