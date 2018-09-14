package http

import (
	"net/http"
	"strconv"

	"sirclo-test/ongkir-app/models"
	OngkirModule "sirclo-test/ongkir-app/ongkir"

	"github.com/labstack/echo"
)

// OngkirHandler is an object for wrapping delivery method
type OngkirHandler struct {
	OngkirUC OngkirModule.OngkirUsecase
}

// OngkirRequest is an object helper to populate request data
type OngkirRequest struct {
	Weight   float64 `form:"weight"`
	Provider int     `form:"provider"`
}

// OngkirResponse is an object helper to populate response data
type OngkirResponse struct {
	Weight   float64
	Provider int
	Result   float64
	Ongkirs  []models.OngkirType
}

// Index function handle Index Page
func (oh *OngkirHandler) Index(c echo.Context) error {
	var weight float64
	var provider int

	// populate ongkir
	ongkirs := []models.OngkirType{
		models.OngkirType{
			OngkirName: "REGULER",
			OngkirType: 0,
		},
		models.OngkirType{
			OngkirName: "EXPRESS",
			OngkirType: 1,
		},
		models.OngkirType{
			OngkirName: "INSTANT",
			OngkirType: 2,
		},
	}

	// get weight cookie
	weightCookie, err := c.Cookie("weight")
	if weightCookie != nil && err == nil {
		weight, err = strconv.ParseFloat(weightCookie.Value, 64)
		if err != nil {
			weight = 0.0
		}
		// delete old cookie
		weightCookie.MaxAge = -1
		c.SetCookie(weightCookie)
	}
	// get provider cookie
	providerCookie, err := c.Cookie("provider")
	if providerCookie != nil && err == nil {
		provider, err = strconv.Atoi(providerCookie.Value)
		if err != nil {
			provider = 0
		}
		// delete old cookie
		providerCookie.MaxAge = -1
		c.SetCookie(providerCookie)
	}

	// calculate
	result := oh.OngkirUC.Calculate(provider, weight)
	return c.Render(http.StatusOK, "ongkir_index", &OngkirResponse{
		Weight:   weight,
		Provider: provider,
		Result:   result,
		Ongkirs:  ongkirs,
	})

}

// Calculate function handle Calculate Request
func (oh *OngkirHandler) Calculate(c echo.Context) error {
	// parse data to model
	var data OngkirRequest
	err := c.Bind(&data)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	// set cookie with request data and redirect to "/"
	SetRequestFlush(c, data.Weight, data.Provider)
	return c.Redirect(http.StatusSeeOther, "/")
}

// SetRequestFlush is Helper Function to store the request in the cookies
func SetRequestFlush(c echo.Context, weight float64, provider int) {
	c.SetCookie(&http.Cookie{
		Name:  "weight",
		Value: strconv.FormatFloat(weight, 'f', 2, 64),
	})
	c.SetCookie(&http.Cookie{
		Name:  "provider",
		Value: strconv.Itoa(provider),
	})
}

// NewOngkirHandler act like constructor of handler
func NewOngkirHandler(e *echo.Echo, uc OngkirModule.OngkirUsecase) {
	handler := &OngkirHandler{
		OngkirUC: uc,
	}
	e.GET("/", handler.Index)
	e.POST("/", handler.Calculate)
}
