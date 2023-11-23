package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"smashedbits.com/shorty/pkg/services"
)

func AuthenticationPage() echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		return eCtx.Render(http.StatusOK, "pages/auth.html", map[string]interface{}{})
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
