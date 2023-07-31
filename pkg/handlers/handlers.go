package handlers

import (
	"net/http"

	"github.com/varunkverma/bedAndBreakfast/pkg/config"
	"github.com/varunkverma/bedAndBreakfast/pkg/models"
	"github.com/varunkverma/bedAndBreakfast/pkg/render"
)

// Repository type to hold dependencies
type Repository struct {
	AppConfig *config.AppConfig
}

// Repo is the respository to which handlers are gonna be linked to
var Repo *Repository

// NewRepository creates a new repository
func NewRepository(appConfig *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: appConfig,
	}
}

// InitHandlersRepository initializes the repository
func InitHandlersRepository(r *Repository) {
	Repo = r
}

// Home is the Home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repo.AppConfig.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the About page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string, 0)
	stringMap["test"] = "hello"

	remoteIp := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	// send the data to the template and render it
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
