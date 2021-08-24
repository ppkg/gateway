package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Any("/*", func(c echo.Context) error {
		//
		//resBody := c.Request().Body
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(resBody)
		//fmt.Println(buf.String())
		return c.String(http.StatusOK, "bbbb")
	})

	e.Start(":8080")
}
