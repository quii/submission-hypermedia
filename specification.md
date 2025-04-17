# Submission Hypermedia Demo - Specification

## 1. Project Structure
submission-hypermedia/
├── main.go              # Application entry point
├── server.go            # HTTP server setup
├── handlers.go          # HTTP route handlers
├── templates.go         # HTML template loading and rendering
├── models.go            # Data models (Submission)
├── static/              # Static assets
│   ├── css/             # CSS files
│   │   └── styles.css   # Basic styling
│   └── js/              # JavaScript files
│       └── htmx.min.js  # HTMX library
└── templates/           # HTML templates
├── layout.html      # Base layout with navigation
├── pages/           # Page templates
│   ├── page1.html   # First dummy page
│   ├── page2.html   # Second dummy page
│   └── page3.html   # Third dummy page
└── components/      # Reusable components
└── submission.html  # Submission hypermedia component


## 2. Data Models

### Submission
```go
type Submission struct {
Title       string
Abstract    string
Keywords    []string
Author      string
Status      string // "pending", "approved", "rejected"
OnWatchlist bool
}
```



## 3. Pages and Navigation
The application will have three dummy pages:

- Submission Overview - Basic information display
- Review Form - A form for reviewing the submission
- Editor Notes - A page for adding notes to the submission

Each page will have:

- The common navigation menu
- A main content area with page-specific dummy content
- The embedded submission hypermedia component

## 4. Submission Hypermedia Component
The submission component will:

- Display the submission details (title, abstract, keywords, author)
- Show the current status (pending, approved, rejected)
- Indicate if it's on the watchlist

Include controls to:
- Approve the submission
- Reject the submission
- Add/remove from watchlist

This component will be loaded via HTMX on each page and will update in-place when actions are performed.

## 5. API Endpoints
GET  /                           # Redirect to first page
GET  /pages/{page-name}          # Render a specific page
GET  /components/submission      # Get the submission component HTML
POST /submission/approve         # Approve the submission
POST /submission/reject          # Reject the submission
POST /submission/toggle-watchlist # Toggle watchlist status



## 6. Implementation Details
### Server
Simple Go HTTP server using the standard library
In-memory storage of a single submission instance
No persistence between server restarts
### Templates
Go's html/template for rendering
Partials for the submission component
Base layout with navigation
### HTMX Integration
Load the submission component with hx-get
Update the submission with hx-post and hx-swap
No full page reloads for submission actions
### Styling
Basic CSS3 for layout and visual hierarchy
Simple visual indicators for submission status
Minimal styling for the navigation and forms
## 7. User Flow
User navigates to any of the three pages
Each page loads with its specific content and the shared submission component
User can interact with the submission controls on any page
When a control is activated, only the submission component updates
The updated state is reflected on all pages when navigating between them
