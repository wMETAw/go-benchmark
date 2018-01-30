package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	// hello worldをレスポンス
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"Hello World")
	})

	// json形式でレスポンス
	e.GET("/users/:id", func(c echo.Context) error {
		jsonMap := map[string]string{
			"id": c.Param("id"),
			"name":"Yamada",
		}
		return c.JSON(http.StatusOK,jsonMap)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
