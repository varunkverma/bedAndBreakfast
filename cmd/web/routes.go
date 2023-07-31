package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/varunkverma/bedAndBreakfast/pkg/config"
	"github.com/varunkverma/bedAndBreakfast/pkg/handlers"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// register routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
