package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/services"
)

func Landing(renderer services.Renderer, auth services.Authenticator, shortener services.Shortener) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		req := eCtx.Request()
		ctx := req.Context()

		user, err := auth.GetUser(eCtx)
		if err != nil {
			eCtx.Logger().Debug(err)
		}

		urls, err := shortener.GetUserURLs(ctx, user.ID)
		if err != nil {
			eCtx.Logger().Error(err)
		}

		tmplVars := services.TemplateVars{
			User: user,
			Data: urls,
		}

		if err := renderer.Render(eCtx, http.StatusOK, "landing.go.html", tmplVars); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return nil
	}
}
