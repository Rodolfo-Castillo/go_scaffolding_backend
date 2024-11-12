package example

import (
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo){
	e.GET("/", func(c echo.Context) error {
		return Hello(c)
	})
}