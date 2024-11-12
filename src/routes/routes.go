package routes

import (
	"myapp/src/modules/example"

	"github.com/labstack/echo/v4"
)

// SetupRoutes configura todas las rutas
func GetAllRoutes(e *echo.Echo) {
	// Definir rutas
	example.SetUpRoutes(e)
}