package main

import (
	"github.com/madecanggih/travel-planner-backend/models"
	"github.com/madecanggih/travel-planner-backend/routes"
)

func main() {
	models.InitDB()
	defer models.CloseDB()

	e := routes.InitRoutes()

	e.Logger.Fatal(e.Start(":8000"))
}
