package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "assets")
	initRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

