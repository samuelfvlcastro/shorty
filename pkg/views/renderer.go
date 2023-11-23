package views

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

const (
	pagesPath      = "pages"
	componentsPath = "components"

	layoutsParts    = "layouts/*.html"
	componentsParts = "components/*/*.html"
)

//go:embed pages/* layouts/* components/**/*
var files embed.FS

type Renderer struct {
	domain    string
	templates map[string]*template.Template
}

func NewRenderer(domain string) *Renderer {
	r := &Renderer{
		domain:    domain,
		templates: make(map[string]*template.Template),
	}
	r.loadTemplates()

	return r
}

func (r *Renderer) loadTemplates() error {
	if err := r.load(pagesPath); err != nil {
		return err
	}

	if err := r.load(componentsPath); err != nil {
		return err
	}

	return nil
}

func (r *Renderer) load(path string) error {
	dir, err := fs.ReadDir(files, path)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		if entry.IsDir() {
			r.load(path + "/" + entry.Name())
			continue
		}

		pt, err := template.ParseFS(files, path+"/"+entry.Name(), layoutsParts, componentsParts)
		if err != nil {
			return err
		}

		r.templates[path+"/"+entry.Name()] = pt
	}
	return nil
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, eCtx echo.Context) error {
	tmpl, ok := r.templates[name]
	if !ok {
		err := errors.New("Template not found:" + name)
		return err
	}

	tmplName := r.fetchTemplateName(name)
	err := tmpl.Funcs(template.FuncMap{
		"composeURL": func(hash string) string {
			return fmt.Sprintf("%s/%s", r.domain, hash)
		},
	}).ExecuteTemplate(w, tmplName, data)
	eCtx.Logger().Info(err)

	return nil
}

func (r *Renderer) fetchTemplateName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		return parts[1]
	}

	return parts[0]
}
