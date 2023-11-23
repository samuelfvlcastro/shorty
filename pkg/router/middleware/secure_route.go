package middleware

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/services"
)

type SecureRouteConfig struct {
	auth services.Authenticator
}

func SecureRoute(auth services.Authenticator) echo.MiddlewareFunc {
	return SecureRouteWithConfig(SecureRouteConfig{
		auth: auth,
	})
}

func AddUserIdToCtx(auth services.Authenticator) echo.MiddlewareFunc {
	return AddUserIdToCtxWithConfig(SecureRouteConfig{
		auth: auth,
	})
}

func SecureRouteWithConfig(config SecureRouteConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) (err error) {
			userId, err := config.auth.GetUserID(eCtx)
			if err != nil {
				htmxAuthRedirect(eCtx)
				return eCtx.JSON(http.StatusUnauthorized, nil)
			}
			eCtx.Set("userId", userId)
			return next(eCtx)
		}
	}
}

func AddUserIdToCtxWithConfig(config SecureRouteConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) (err error) {
			userId, err := config.auth.GetUserID(eCtx)
			if err != nil {
				return next(eCtx)
			}
			eCtx.Set("userId", userId)
			return next(eCtx)
		}
	}
}

func hashQueryString(req *http.Request) []byte {
	qs := req.URL.Query().Encode()

	hash := md5.New()
	io.WriteString(hash, qs)

	return hash.Sum(nil)
}

func htmxAuthRedirect(eCtx echo.Context) {
	hash := hashQueryString(eCtx.Request())
	redirectData := fmt.Sprintf("/auth?hash=%s", hash)
	eCtx.Response().Header().Add("HX-Redirect", redirectData)
}
