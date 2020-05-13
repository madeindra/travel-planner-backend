package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/madecanggih/travel-planner-backend/handlers"
	"github.com/madecanggih/travel-planner-backend/helpers"
	"github.com/madecanggih/travel-planner-backend/models"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	ah := handlers.NewAuthHandler(models.NewAuthImplementation(models.DB), helpers.NewHelperImplementation())
	uh := handlers.NewUsersHandler(models.NewUsersImplementation(models.DB), helpers.NewValidatorImplementation())

	api := e.Group("/api/v1")
	{
		api.GET("/users", uh.GetAllUsers)
		api.GET("/users/:id", uh.GetOneUser)

		api.POST("/auth/authentication", ah.PostLogin)
		api.POST("/auth/registration", ah.PostRegister)
	}

	return e
}
