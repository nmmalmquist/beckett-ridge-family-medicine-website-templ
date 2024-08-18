package main

import (
	"app/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func indexHandler(ctx echo.Context) error {
	element := pages.Hello("Lets go")
	return render(ctx, element)
}