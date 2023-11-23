package middleware

import "github.com/labstack/echo/v4"

type GothProviderPrefillConfig struct {
}

var DefaultGothProviderPrefillConfig = GothProviderPrefillConfig{}

func GothProviderPrefill() echo.MiddlewareFunc {
	return GothProviderPrefillWithConfig(DefaultGothProviderPrefillConfig)
}

func GothProviderPrefillWithConfig(config GothProviderPrefillConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) (err error) {
			if provider := eCtx.Param("provider"); provider != "" {
				qs := eCtx.Request().URL.Query()
				qs.Add("provider", provider)
				eCtx.Request().URL.RawQuery = qs.Encode()
			}
			return next(eCtx)
		}
	}
}
