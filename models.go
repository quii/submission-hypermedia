package main

// Submission represents a paper submission to a peer review system
type Submission struct {
	Title       string
	Abstract    string
	Keywords    []string
	Author      string
	Status      string // "pending", "approved", "rejected"
	OnWatchlist bool
}

// DefaultSubmission creates a sample submission with default values
func DefaultSubmission() *Submission {
	return &Submission{
		Title:       "Advances in Hypermedia-Driven Web Applications",
		Abstract:    "This paper explores how hypermedia controls can simplify web application architecture while improving user experience and maintainability.",
		Keywords:    []string{"hypermedia", "web architecture", "HATEOAS", "user experience"},
		Author:      "Jane Researcher",
		Status:      "pending",
		OnWatchlist: false,
	}
}
