package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"smashedbits.com/shorty/pkg/services"
)

func AuthenticationPage(renderer services.Renderer, auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		user, err := auth.GetUser(eCtx)
		if err != nil {
			eCtx.Logger().Debug(err)
		}

		tmplVars := services.TemplateVars{
			User: user,
		}

		if err := renderer.Render(eCtx, http.StatusOK, "authentication.go.html", tmplVars); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return nil
	}
}

func AuthProvider(auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		if err := auth.CompleteUserSignIn(eCtx, true); err != nil {
			eCtx.Logger().Debug(err)
			return nil
		}

		return eCtx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func AuthProviderCallback(auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		if err := auth.CompleteUserSignIn(eCtx, false); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not complete user auth")
		}

		return eCtx.Redirect(http.StatusMovedPermanently, "/")
	}
}

func AuthProviderLogout(auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		if err := auth.Logout(eCtx); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not logout user")
		}

		return eCtx.Redirect(http.StatusMovedPermanently, "/")
	}
}
