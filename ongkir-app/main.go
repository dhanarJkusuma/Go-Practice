package main

import (
	OngkirModuleHandler "sirclo-test/ongkir-app/ongkir/delivery/http"
	OngkirModuleUsecase "sirclo-test/ongkir-app/ongkir/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	ou := OngkirModuleUsecase.NewOngkirUsecase()
	OngkirModuleHandler.NewOngkirHandler(e, ou)

	SetRenderTemplate(e)

	SetCustomNotFoundHandler()

	e.Logger.Fatal(e.Start(":8080"))
}
