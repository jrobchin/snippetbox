package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func (app *application) renderTemplate(w http.ResponseWriter, templatePath string) error {
	files := []string{
		filepath.Join("../../ui/html", templatePath),
		filepath.Join("../../ui/html/footer.partial.tmpl"),
		filepath.Join("../../ui/html/base.layout.tmpl"),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return err
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}

	return nil
}

// Home handler that writes a hello message
// Define methods on the application struct to gain access to dependencies
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	app.renderTemplate(w, "home.page.tmpl")
}

// Returns a specified snipped
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snipped with ID %d...", id)
}

// Creates a snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a snippet"))
}
