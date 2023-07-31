package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/varunkverma/bedAndBreakfast/pkg/config"
	"github.com/varunkverma/bedAndBreakfast/pkg/models"
)

var localConfig *config.AppConfig

// NewTemplates sets the config for the render package
func NewTemplates(appConfig *config.AppConfig) {
	localConfig = appConfig
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// CreateTemplateCache parses the pages and layouts present in ./templates folder and creates a map of file name to template
func CreateTemplateCache() (map[string]*template.Template, error) {
	log.Println("Creating Template Cache")

	tmplCache := make(map[string]*template.Template, 0)

	// getting file path of the page-template files, which ends with .page.tmpl
	pageFilePaths, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	// getting file path of the layout-template files, which ends with .layout.tmpl
	layoutFilePaths, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return tmplCache, err
	}

	// parsing
	for _, pagePath := range pageFilePaths {
		pageName := filepath.Base(pagePath)

		// parseing the page template, to create a template with page's name
		parsedTemplate, err := template.New(pageName).ParseFiles(pagePath)
		if err != nil {
			log.Println("error parsing template", pageName, err.Error())
			return tmplCache, err
		}

		//parsing the layout templates, if any
		if len(layoutFilePaths) > 0 {
			parsedTemplate, err = parsedTemplate.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("error parsing template with layouts", pageName, err.Error())
				return tmplCache, err
			}
		}

		tmplCache[pageName] = parsedTemplate
	}
	return tmplCache, nil
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template
	var err error

	if localConfig.UseCache {
		// get the requested template from the template cache
		templateCache = localConfig.TemplateCache
	} else {
		// create a fresh template cache
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("cannot create template cache dynamically", err.Error())
		}
	}

	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatalf("cannot find %s temaple in the template cache", tmpl)
	}

	templateData = addDefaultData(templateData)

	// create a byte buffer to write the result of template getting executed
	buf := new(bytes.Buffer)

	err = template.Execute(buf, templateData)
	if err != nil {
		log.Println("error while executing tempplate", err.Error())
	}

	// writing the result of template execution from the buffer to the response writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error parsing template", tmpl, err.Error())
	}
}
