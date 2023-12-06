package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/services"
)

func ToggleDarkMode(renderer services.Renderer, auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		user, err := auth.GetUser(eCtx)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, nil)
		}

		user.DarkMode = !user.DarkMode
		if err := auth.UpdateUser(eCtx, user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not toggle darkmode")
		}

		tmplVars := services.TemplateVars{
			User: user,
		}

		if err := renderer.Render(eCtx, http.StatusOK, "components/darkmode_toggle.go.html", tmplVars); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return nil
	}
}
