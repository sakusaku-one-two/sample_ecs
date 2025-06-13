package main

import (
	"module/infra"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := infra.FactoryServer()
	e.GET("/", func(c echo.Context) error {
		res := make(map[string]string)
		res["value"] = "hello my server"
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(infra.SERVER_PORT))
}
