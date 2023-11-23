package middleware

// source:
// https://github.com/LYY/echo-middleware/blob/master/nocache.go

import (
	"time"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type (
	// NoCacheConfig defines the config for nocache middleware.
	NoCacheConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper emw.Skipper
	}
)

var (
	// Unix epoch time
	epoch = time.Unix(0, 0).Format(time.RFC1123)

	// Taken from https://github.com/mytrile/nocache
	noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	}
	etagHeaders = []string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	}
	// DefaultNoCacheConfig is the default nocache middleware config.
	DefaultNoCacheConfig = NoCacheConfig{
		Skipper: emw.DefaultSkipper,
	}
)

func NoCache() echo.MiddlewareFunc {
	return NoCacheWithConfig(DefaultNoCacheConfig)
}

func NoCacheWithConfig(config NoCacheConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultNoCacheConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) (err error) {
			if config.Skipper(eCtx) {
				return next(eCtx)
			}
			req := eCtx.Request()
			for _, v := range etagHeaders {
				if req.Header.Get(v) != "" {
					req.Header.Del(v)
				}
			}

			res := eCtx.Response()
			for k, v := range noCacheHeaders {
				res.Header().Set(k, v)
			}

			return next(eCtx)
		}
	}
}
