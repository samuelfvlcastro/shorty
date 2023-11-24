package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"smashedbits.com/shorty/pkg/services"
	"smashedbits.com/shorty/pkg/views/layouts"
	"smashedbits.com/shorty/pkg/views/pages"
)

func AuthenticationPage(auth services.Authenticator) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		user, _ := auth.GetUser(eCtx)

		a := &pages.Auth{
			UserIdStg:    user.ID,
			UserEmailStg: user.Email,
		}
		layouts.WriteBaseLayout(eCtx.Response().Writer, a)
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
