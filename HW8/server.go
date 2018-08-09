//寫一個服務server，要有兩個以上的route與一個端程式去call服務

package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/money", getMoney)
	e.GET("/anyththing/:thing", getAnything)

	e.Logger.Fatal(e.Start(":1323"))
}

func getMoney(c echo.Context) error {
	return c.String(http.StatusOK, "xxxx")
}

func getAnything(c echo.Context) error {
	thing := c.Param("thing")
	return c.String(http.StatusOK, thing)
}
