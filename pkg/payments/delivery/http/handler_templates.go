package http

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()
	return tpl
}

// ReloadTemplates method will will parse template
func (r *Renderer) ReloadTemplates() {
	r.template = template.Must(template.ParseGlob(r.location))
}

// Render method is used for rendering templates that has been parsed as an output
func (r *Renderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	if r.debug {
		r.ReloadTemplates()
	}

	return r.template.ExecuteTemplate(w, name, data)
}
