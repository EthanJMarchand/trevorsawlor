package handlers

import (
	"github.com/ethanjmarchand/trevorsawlor/internal/render"
	"net/http"
)

// Home is the home handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.gohtml")
}

// About is the handler for the about page route.
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.gohtml")
}
