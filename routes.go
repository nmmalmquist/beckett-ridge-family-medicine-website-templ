package main

import "github.com/labstack/echo"


func initRoutes (e *echo.Echo) {
	e.GET("/", indexHandler)
}