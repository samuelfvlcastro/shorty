package router

import (
	"github.com/labstack/echo/v4"
	"smashedbits.com/shorty/pkg/handlers"
	"smashedbits.com/shorty/pkg/router/middleware"
	"smashedbits.com/shorty/pkg/services"
)

func InitRoutes(app *echo.Echo, auth services.Authenticator, shortener services.Shortener) {
	app.Static("/dist", "ui/dist")
	app.Static("/css", "ui/src/css")

	mainRoutes := app.Group("")
	mainRoutes.Use(middleware.AddUserIdToCtx(auth))
	mainRoutes.GET("/", handlers.Landing(auth, shortener))

	secureRoutes := app.Group("")
	secureRoutes.Use(middleware.SecureRoute(auth))
	secureRoutes.POST("/url", handlers.InsertURL(auth, shortener))

	authRoutes := app.Group("/auth")
	authRoutes.Use(middleware.GothProviderPrefill())
	authRoutes.Use(middleware.NoCache())

	authRoutes.GET("", handlers.AuthenticationPage())
	authRoutes.GET("/:provider", handlers.AuthProvider(auth))
	authRoutes.GET("/:provider/callback", handlers.AuthProviderCallback(auth))
	authRoutes.GET("/logout", handlers.AuthProviderLogout(auth))
}
