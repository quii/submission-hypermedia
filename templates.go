package main

import (
	"html/template"
	"log"
	"path/filepath"
)

// Templates holds all parsed HTML templates
type Templates struct {
	layout     *template.Template
	pages      map[string]*template.Template
	components map[string]*template.Template
}

// LoadTemplates loads and parses all HTML templates
func LoadTemplates() *Templates {
	t := &Templates{
		pages:      make(map[string]*template.Template),
		components: make(map[string]*template.Template),
	}

	// Parse base layout
	layoutPath := filepath.Join("templates", "layout.html")
	layoutTmpl, err := template.ParseFiles(layoutPath)
	if err != nil {
		log.Fatalf("Error parsing layout template: %v", err)
	}
	t.layout = layoutTmpl

	// Parse page templates
	pageFiles := []string{"page1.html", "page2.html", "page3.html"}
	for _, file := range pageFiles {
		pagePath := filepath.Join("templates", "pages", file)
		// Clone the layout template and add the page template
		pageTmpl, err := t.layout.Clone()
		if err != nil {
			log.Fatalf("Error cloning layout template: %v", err)
		}
		pageTmpl, err = pageTmpl.ParseFiles(pagePath)
		if err != nil {
			log.Fatalf("Error parsing page template %s: %v", file, err)
		}
		// Store with name without extension (e.g., "page1")
		pageName := file[:len(file)-5]
		t.pages[pageName] = pageTmpl
	}

	// Parse component templates
	componentFiles := []string{"submission.html"}
	for _, file := range componentFiles {
		componentPath := filepath.Join("templates", "components", file)
		componentTmpl, err := template.ParseFiles(componentPath)
		if err != nil {
			log.Fatalf("Error parsing component template %s: %v", file, err)
		}
		// Store with name without extension (e.g., "submission")
		componentName := file[:len(file)-5]
		t.components[componentName] = componentTmpl
	}

	return t
}
