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
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return err
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	return nil
}

// Home handler that writes a hello message
// Define methods on the application struct to gain access to dependencies
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.renderTemplate(w, "home.page.tmpl")
}

// Returns a specified snipped
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snipped with ID %d...", id)
}

// Creates a snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("create a snippet"))
}
