package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	IsProduction  bool
	UseCache      bool
	Session       *scs.SessionManager
}
