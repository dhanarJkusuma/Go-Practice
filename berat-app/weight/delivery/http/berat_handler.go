package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"sirclo-test/berat-app/models"
	WeightModule "sirclo-test/berat-app/weight"
)

type WeightHandler struct {
	WeightUC WeightModule.WeightUsecase
}

func (w *WeightHandler) Index(c echo.Context) error {
	data := w.WeightUC.ListRecord()
	return c.Render(http.StatusOK, "weight_index", data)
}

func (w *WeightHandler) CreateForm(c echo.Context) error {
	return c.Render(http.StatusOK, "weight_create", nil)
}

func (w *WeightHandler) Create(c echo.Context) error {
	var data models.Weight
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, "weight_create", "Error !")
	}

	// parse Date
	t, err := time.Parse("2006-01-02", data.RequestDate)
	if err != nil {
		return c.Render(http.StatusOK, "weight_create", "Error, Data `Date` salah format, harus `yyyy-mm-dd` !")
	}
	data.Date = t
	isOk, msg := w.WeightUC.CreateRecord(&data)
	if isOk {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return c.Render(http.StatusOK, "weight_create", msg)
}

func (w *WeightHandler) UpdateForm(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		//return c.Render(http.StatusNotFound, "/", nil)
		return c.Redirect(http.StatusSeeOther, "/")
	}
	isOk, data := w.WeightUC.DetailRecord(idInt)
	if !isOk {
		// return 404
	}
	return c.Render(http.StatusOK, "weight_update", data)
}

func (w *WeightHandler) Update(c echo.Context) error {
	var data models.Weight

	// validate param id
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		//return c.Render(http.StatusNotFound, "/", nil)
		return c.Redirect(http.StatusSeeOther, "/")
	}

	err = c.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, "weight_update", "Error !")
	}

	// parse Date
	t, err := time.Parse("2006-01-02", data.RequestDate)
	if err != nil {
		return c.Render(http.StatusOK, "weight_update", "Error, Data `Date` salah format, harus `yyyy-mm-dd` !")
	}
	data.Date = t
	isOk, msg := w.WeightUC.UpdateRecord(idInt, data)
	if isOk {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return c.Render(http.StatusOK, "weight_update", msg)
}

func NewWeightHandler(e *echo.Echo, uw WeightModule.WeightUsecase) {
	handler := WeightHandler{
		WeightUC: uw,
	}
	e.GET("/", handler.Index)
	e.GET("/create", handler.CreateForm)
	e.POST("/create", handler.Create)
	e.GET("/update/:id", handler.UpdateForm)
	e.POST("/update/:id", handler.Update)
}
