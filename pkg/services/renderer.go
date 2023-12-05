package services

import (
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/model"
	"smashedbits.com/shorty/pkg/templates"
)

type Page interface {
	User() model.User
}

type TemplateVars struct {
	User model.User
	Data interface{}
}

const HeaderContentType = "Content-Type"
const ContentTypeTextHTML string = "text/html; charset=UTF-8"

type Renderer struct {
	tmpls map[string]*template.Template
}

func NewRenderer() (Renderer, error) {
	funcMap := template.FuncMap{
		"slice_first_and_up": func(str string) string {
			return strings.ToUpper(str[0:1])
		},
	}

	tmpls := map[string]*template.Template{}

	layoutFiles := []string{}
	if err := fs.WalkDir(templates.FS, "layouts", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		layoutFiles = append(layoutFiles, path)
		return nil
	}); err != nil {
		return Renderer{}, err
	}

	componentsFiles := []string{}
	if err := fs.WalkDir(templates.FS, "components", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		componentsFiles = append(componentsFiles, path)
		return nil
	}); err != nil {
		return Renderer{}, err
	}

	if err := fs.WalkDir(templates.FS, "pages", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		fileName := filepath.Base(path)
		files := append(layoutFiles, componentsFiles...)
		files = append(files, path)

		tmpls[fileName] = template.Must(template.New(fileName).Funcs(funcMap).ParseFS(templates.FS, files...))

		return nil
	}); err != nil {
		return Renderer{}, err
	}

	return Renderer{
		tmpls: tmpls,
	}, nil
}

func (r *Renderer) Render(eCtx echo.Context, httpCode int, tmplName string, data TemplateVars) error {
	w := eCtx.Response().Writer

	r.writeContentType(ContentTypeTextHTML, eCtx)
	eCtx.Response().WriteHeader(httpCode)

	return r.tmpls[tmplName].ExecuteTemplate(w, tmplName, data)
}

func (r *Renderer) writeContentType(value string, eCtx echo.Context) {
	header := eCtx.Response().Header()
	if header.Get(HeaderContentType) == "" {
		header.Set(HeaderContentType, value)
	}
}
