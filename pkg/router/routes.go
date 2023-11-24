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

	mainRoutes := app.Group("",
		middleware.GothProviderPrefill(),
		middleware.NoCache(),
	)
	mainRoutes.GET("/", handlers.Landing(auth, shortener))
	mainRoutes.GET("/auth", handlers.AuthenticationPage(auth))
	mainRoutes.GET("/auth/:provider", handlers.AuthProvider(auth))
	mainRoutes.GET("/auth/:provider/callback", handlers.AuthProviderCallback(auth))
	mainRoutes.GET("/auth/logout", handlers.AuthProviderLogout(auth))

	secureRoutes := app.Group("", middleware.SecureRoute(auth))
	secureRoutes.POST("/url", handlers.InsertURL(auth, shortener))
}
