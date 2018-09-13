package main

import (
	"html/template"
	"io"
	"sirclo-test/berat-app/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func parseDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func diffWeight(weight models.Weight) string {
	res := weight.WeightMax - weight.WeightMin
	return strconv.FormatFloat(res, 'f', 1, 64)
}

func SetRenderTemplate(e *echo.Echo) {
	var variableFunc = template.FuncMap{
		"parseDate":  parseDate,
		"diffWeight": diffWeight,
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("weight").Funcs(variableFunc).ParseGlob("views/berat/*.gohtml")),
	}
	e.Renderer = renderer
}
