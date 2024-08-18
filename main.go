package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	initRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

