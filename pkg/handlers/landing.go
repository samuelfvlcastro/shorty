package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/services"
	"smashedbits.com/shorty/pkg/views/components/url"
	"smashedbits.com/shorty/pkg/views/layouts"
	"smashedbits.com/shorty/pkg/views/pages"
)

func Landing(auth services.Authenticator, shortener services.Shortener) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		req := eCtx.Request()
		ctx := req.Context()

		user, _ := auth.GetUser(eCtx)
		urls, err := shortener.GetUserURLs(ctx, user.ID)
		if err != nil {
			eCtx.Logger().Error(err)
		}

		p := &pages.Landing{
			UserIdStg:    user.ID,
			UserEmailStg: user.Email,
			Urls:         urls,
		}
		layouts.WriteBaseLayout(eCtx.Response().Writer, p)

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

		eCtx.Response().Writer.Write([]byte(url.RenderList(urls)))
		return nil
	}
}
