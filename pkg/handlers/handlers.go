package handlers

import (
	"net/http"

	"github.com/tanishqv/mywebapp-go/pkg/config"
	"github.com/tanishqv/mywebapp-go/pkg/render"
)

// Handlers may not use template cache, but the config may be updated with things that makes the application run better
// Repository pattern to swap components out of the application with minimal changes erquired to code base

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is thr repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
