package view

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// templ view render with echo framework
func Render(c echo.Context, component templ.Component) error {
	if err := component.Render(c.Request().Context(), c.Response().Writer); err != nil {
		return c.JSON(http.StatusInternalServerError, "Template rendering failed")
	}
	return nil
}
