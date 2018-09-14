package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"sirclo-test/berat-app/models"
	WeightModule "sirclo-test/berat-app/weight"
)

// WeightHandler
type WeightHandler struct {
	WeightUC WeightModule.WeightUsecase
}

// DeleteRequest is a model temporary for getting request id
type DeleteRequest struct {
	ID int64 `form:"id"`
}

// Index Page Handler
func (w *WeightHandler) Index(c echo.Context) error {
	var isErr bool
	var msg string
	var page, size string
	var pg, sz int
	var err error

	// get page and size from query
	if page = c.QueryParam("page"); page == "" {
		page = "0"
	}
	if size = c.QueryParam("size"); size == "" {
		size = "10"
	}

	// parsing to int
	if pg, err = strconv.Atoi(page); err != nil {
		return c.Redirect(http.StatusOK, "/")
	}
	if sz, err = strconv.Atoi(size); err != nil {
		return c.Redirect(http.StatusOK, "/")
	}

	data := w.WeightUC.ListRecord(pg, sz)

	// Set cookie to the variable and clear old cookie
	errCookie, err := c.Cookie("isError")
	if errCookie != nil && err == nil {
		isErr, err = strconv.ParseBool(errCookie.Value)
		errCookie.MaxAge = -1
		c.SetCookie(errCookie)
	}
	msgCookie, err := c.Cookie("msg")
	if msgCookie != nil && err == nil {
		msg = msgCookie.Value
		msgCookie.MaxAge = -1
		c.SetCookie(msgCookie)
	}

	// calculate total pagination
	data.LastPage = data.TotalCount / data.Size
	if data.TotalCount%data.Size > 0 {
		data.LastPage++
	}

	// construct response
	res := &models.CommonResponse{
		Message: msg,
		IsError: isErr,
		Data:    data,
	}

	return c.Render(http.StatusOK, "weight_index", res)
}

// CreateForm Page Handler
func (w *WeightHandler) CreateForm(c echo.Context) error {
	return c.Render(http.StatusOK, "weight_create", nil)
}

// Create Request Page Handler
func (w *WeightHandler) Create(c echo.Context) error {
	var data models.Weight
	err := c.Bind(&data)
	if err != nil {
		return c.Render(http.StatusOK, "weight_create", &models.CommonResponse{
			IsError: true,
			Message: err.Error(),
		})
	}

	// parse Date
	t, err := time.Parse("2006-01-02", data.RequestDate)
	if err != nil {
		return c.Render(http.StatusOK, "weight_create", &models.CommonResponse{
			IsError: true,
			Message: "Error, Data `Date` salah format, harus `yyyy-mm-dd` !",
		})
	}
	data.Date = t

	// create record
	isOk, msg := w.WeightUC.CreateRecord(&data)
	if isOk {
		SetResponseFlushMessage(c, false, msg)
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return c.Render(http.StatusOK, "weight_create", &models.CommonResponse{
		IsError: true,
		Message: msg,
	})
}

// UpdateForm Page Handler
func (w *WeightHandler) UpdateForm(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// return 404
		return echo.NotFoundHandler(c)
	}
	isOk, data := w.WeightUC.DetailRecord(idInt)
	if !isOk {
		// return 404
		return echo.NotFoundHandler(c)
	}
	return c.Render(http.StatusOK, "weight_update", &models.CommonResponse{
		IsError: false,
		Message: "",
		Data:    data,
	})
}

// Update Request Page Handler
func (w *WeightHandler) Update(c echo.Context) error {
	var data models.Weight

	// validate param id
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	// get existing data
	isOk, existData := w.WeightUC.DetailRecord(idInt)
	if !isOk {
		return echo.NotFoundHandler(c)
	}

	// bind data
	err = c.Bind(&data)
	if err != nil {
		return c.Render(http.StatusOK, "weight_update", &models.CommonResponse{
			IsError: true,
			Message: err.Error(),
			Data:    existData,
		})
	}

	// parse Date
	t, err := time.Parse("2006-01-02", data.RequestDate)
	if err != nil {
		return c.Render(http.StatusOK, "weight_update", &models.CommonResponse{
			IsError: true,
			Message: "Error, Data `Date` salah format, harus `yyyy-mm-dd` !",
			Data:    data,
		})
	}
	data.Date = t
	isOk, msg := w.WeightUC.UpdateRecord(idInt, data)
	if isOk {
		// set message in cookies
		SetResponseFlushMessage(c, false, "Berhasil mengubah data.")
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return c.Render(http.StatusOK, "weight_update", &models.CommonResponse{
		IsError: true,
		Message: msg,
		Data:    data,
	})
}

// Delete Request Page Handler
func (w *WeightHandler) Delete(c echo.Context) error {
	var req DeleteRequest

	// if error parsing form to model
	err := c.Bind(&req)
	if err != nil {
		// return 404
		return echo.NotFoundHandler(c)
	}

	// passing message response using cookies
	isOk, msg := w.WeightUC.DeleteRecord(req.ID)

	// set message in cookies
	SetResponseFlushMessage(c, !isOk, msg)
	return c.Redirect(http.StatusSeeOther, "/")
}

// Detail Page Handler
func (w *WeightHandler) Detail(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	isOk, data := w.WeightUC.DetailRecord(idInt)
	if !isOk {
		return echo.NotFoundHandler(c)
	}
	return c.Render(http.StatusOK, "weight_detail", data)
}

// SetResponseFlushMessage is Helper Function to Make Message Flush using Cookie
func SetResponseFlushMessage(c echo.Context, isErr bool, msg string) {
	c.SetCookie(&http.Cookie{
		Name:  "isError",
		Value: strconv.FormatBool(isErr),
	})
	c.SetCookie(&http.Cookie{
		Name:  "msg",
		Value: msg,
	})
}

// NewWeightHandler act like constructor
func NewWeightHandler(e *echo.Echo, uw WeightModule.WeightUsecase) {
	handler := WeightHandler{
		WeightUC: uw,
	}
	e.GET("/", handler.Index)
	e.GET("/create", handler.CreateForm)
	e.POST("/create", handler.Create)
	e.GET("/detail/:id", handler.Detail)
	e.GET("/update/:id", handler.UpdateForm)
	e.POST("/update/:id", handler.Update)
	e.POST("/delete", handler.Delete)
}
