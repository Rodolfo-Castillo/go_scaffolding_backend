package main

import (
	"myapp/src/routes"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	routes.GetAllRoutes(e)
	// e.Use(echojwt.JWT([]byte("secret")))
	e.Logger.Fatal(e.Start(":3000"))
}


// func show(c echo.Context) error {
// 	// Get team and member from the query string
// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")
// 	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
// }


