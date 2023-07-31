package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/varunkverma/bedAndBreakfast/pkg/config"
	"github.com/varunkverma/bedAndBreakfast/pkg/handlers"
	"github.com/varunkverma/bedAndBreakfast/pkg/render"
)

const PORT_NUMBER = ":3000"

var appConfig config.AppConfig

var session *scs.SessionManager

func main() {

	appConfig.IsProduction = false // only for dev

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsProduction

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err.Error())
	}

	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = false // only in dev
	appConfig.Session = session

	handlersRepo := handlers.NewRepository(&appConfig)
	handlers.InitHandlersRepository(handlersRepo)

	render.NewTemplates(&appConfig)

	log.Printf("Starting app @ %s\n", PORT_NUMBER)
	routesToRegister := routes(&appConfig)

	server := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routesToRegister,
	}
	err = server.ListenAndServe()
	log.Println(err.Error())
}
