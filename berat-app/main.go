package main

import (
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	weightModuleHandler "sirclo-test/berat-app/weight/delivery/http"
	weightModuleRepo "sirclo-test/berat-app/weight/repository"
	weightModuleUsecase "sirclo-test/berat-app/weight/usecase"
)

func main() {
	// connect db
	db, err := gorm.Open("sqlite3", "db/app.db")
	if err != nil {
		panic("[Weight App] : Failed to connect Database")
	}
	defer db.Close()

	// db migration
	AutoMigration(db)

	// init echo
	e := echo.New()

	// set render template
	SetRenderTemplate(e)

	// connect the module
	wr := weightModuleRepo.NewWeightRepository(db)
	wu := weightModuleUsecase.NewWeightUsecase(wr)
	weightModuleHandler.NewWeightHandler(e, wu)

	// handling not found
	SetCustomNotFoundHandler()

	e.Logger.Fatal(e.Start(":8080"))
}
