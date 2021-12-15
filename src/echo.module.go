package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	//instanciando echo
	myEcho := echo.New()

	//ruta
	myEcho.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world from an API built in Go!")
	})
	myEcho.Logger.Fatal(myEcho.Start(":1323"))
}
