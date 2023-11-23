package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/services"
)

func Landing(auth services.Authenticator, shortener services.Shortener) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		req := eCtx.Request()
		ctx := req.Context()

		userId, _ := eCtx.Get("userId").(string)
		urls, err := shortener.GetUserURLs(ctx, userId)
		if err != nil {
			eCtx.Logger().Error(err)
		}

		if err := eCtx.Render(http.StatusOK, "pages/landing.html", map[string]interface{}{
			"Urls": urls,
		}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func InsertURL(auth services.Authenticator, shortener services.Shortener) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		req := eCtx.Request()
		ctx := req.Context()

		userId, ok := eCtx.Get("userId").(string)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
		}

		longURL := eCtx.FormValue("long_url")
		if longURL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing long_url")
		}

		if _, err := shortener.InsertURL(ctx, userId, longURL); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		urls, err := shortener.GetUserURLs(ctx, userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
		}

		return eCtx.Render(http.StatusOK, "components/url/list", map[string]interface{}{
			"Urls": urls,
		})
	}
}
