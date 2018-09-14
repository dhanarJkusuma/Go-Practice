package main

import (
	"fmt"
	"html/template"
	"io"

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

func parseFloat64(data float64) string {
	return fmt.Sprintf("%.0f", data)
}

func parseToCurrency(data float64) string {
	resS := parseFloat64(data)
	var result string
	var index int
	for i := len(resS) - 1; i >= 0; i-- {
		result += string(resS[i])
		if (index+1)%3 == 0 && i > 0 {
			result += ","
		}
		index++
	}
	return reverseString(result)
}

func reverseString(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

// SetRenderTemplate is function to set rendering source template with Helper Function
func SetRenderTemplate(e *echo.Echo) {
	var variableFunc = template.FuncMap{
		"parseFloat64":    parseFloat64,
		"parseToCurrency": parseToCurrency,
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("OngkirTemplate").Funcs(variableFunc).ParseGlob("views/*.gohtml")),
	}
	e.Renderer = renderer
}
