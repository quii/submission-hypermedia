package main

import (
	"net/http"
	"strings"
	"sync"
)

// AppHandler holds application state and handler methods
type AppHandler struct {
	templates  *Templates
	submission *Submission
	mu         sync.Mutex // Protects submission
}

// NewAppHandler creates a new AppHandler with initialized state
func NewAppHandler() *AppHandler {
	return &AppHandler{
		templates:  LoadTemplates(),
		submission: DefaultSubmission(),
	}
}

// IndexHandler redirects to the first page
func (h *AppHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/pages/page1", http.StatusSeeOther)
}

// PageHandler renders a specific page with the submission component
func (h *AppHandler) PageHandler(w http.ResponseWriter, r *http.Request) {
	pageName := strings.TrimPrefix(r.URL.Path, "/pages/")

	// Check if the page exists
	tmpl, ok := h.templates.pages[pageName]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Render the page template
	h.mu.Lock()
	data := map[string]interface{}{
		"Submission": h.submission,
		"PageName":   pageName,
	}
	h.mu.Unlock()

	tmpl.Execute(w, data)
}

// SubmissionComponentHandler renders just the submission component
func (h *AppHandler) SubmissionComponentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, ok := h.templates.components["submission"]
	if !ok {
		http.NotFound(w, r)
		return
	}

	h.mu.Lock()
	tmpl.Execute(w, h.submission)
	h.mu.Unlock()
}

// ApproveSubmissionHandler approves the submission
func (h *AppHandler) ApproveSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.submission.Status = "approved"
	h.mu.Unlock()

	h.SubmissionComponentHandler(w, r)
}

// RejectSubmissionHandler rejects the submission
func (h *AppHandler) RejectSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.submission.Status = "rejected"
	h.mu.Unlock()

	h.SubmissionComponentHandler(w, r)
}

// ToggleWatchlistHandler toggles the watchlist status
func (h *AppHandler) ToggleWatchlistHandler(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.submission.OnWatchlist = !h.submission.OnWatchlist
	h.mu.Unlock()

	h.SubmissionComponentHandler(w, r)
}
