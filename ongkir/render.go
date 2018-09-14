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
	return t.templates.ExecuteTemplate(w, name, data)
}

func parseDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func diffWeight(weight models.Weight) string {
	res := weight.WeightMax - weight.WeightMin
	return strconv.FormatFloat(res, 'f', 1, 64)
}

func iterate(count int) []int {
	var i int
	var Items []int
	for i = 0; i < count; i++ {
		Items = append(Items, i)
	}
	return Items
}

func inc(count int) string {
	return strconv.Itoa(count + 1)
}

func dec(count int) string {
	return strconv.Itoa(count - 1)
}

// SetRenderTemplate is function to set rendering source template with Helper Function
func SetRenderTemplate(e *echo.Echo) {
	var variableFunc = template.FuncMap{
		"parseDate":  parseDate,
		"diffWeight": diffWeight,
		"iterate":    iterate,
		"inc":        inc,
		"dec":        dec,
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("weight").Funcs(variableFunc).ParseGlob("views/*.gohtml")),
	}
	e.Renderer = renderer
}
