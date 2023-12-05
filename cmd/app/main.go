package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"smashedbits.com/shorty/pkg/router"
	"smashedbits.com/shorty/pkg/services"
	"smashedbits.com/shorty/pkg/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	app := &cli.App{
		Name:        "Shorty",
		Usage:       "simpler link shortener",
		HelpName:    "Shorty",
		Description: "Simpler link shortener for smashedbits tool chain",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "PORT",
				Aliases: []string{"p"},
				Usage:   "Port to run the API in",
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "DOMAIN",
				EnvVars: []string{"DOMAIN"},
			},
			&cli.StringFlag{
				Name:    "DATABASE_URI",
				EnvVars: []string{"DATABASE_URI"},
			},
			&cli.StringFlag{
				Name:    "JWT_SIGNING_KEY",
				EnvVars: []string{"JWT_SIGNING_KEY"},
			},
			&cli.StringFlag{
				Name:    "GOOGLE_KEY",
				EnvVars: []string{"GOOGLE_KEY"},
			},
			&cli.StringFlag{
				Name:    "GOOGLE_SECRET",
				EnvVars: []string{"GOOGLE_SECRET"},
			},
			&cli.StringFlag{
				Name:    "GOOGLE_CALLBACK",
				EnvVars: []string{"GOOGLE_CALLBACK"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return bootstrap(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func bootstrap(ctx *cli.Context) error {
	port := ctx.Int("PORT")
	dbURI := ctx.String("DATABASE_URI")
	jwtSigningKey := ctx.String("JWT_SIGNING_KEY")
	sessionSecret := ctx.String("SESSION_SECRET")
	googleKey := ctx.String("GOOGLE_KEY")
	googleSecret := ctx.String("GOOGLE_SECRET")
	googleCallback := ctx.String("GOOGLE_CALLBACK")

	app := echo.New()
	app.Pre(middleware.RemoveTrailingSlash())

	conn, err := pgx.Connect(ctx.Context, dbURI)
	if err != nil {
		return err
	}
	defer conn.Close(ctx.Context)

	urlStore := storage.NewURLs(conn)
	userStore := storage.NewUsers(conn)

	gothic.Store = sessions.NewCookieStore([]byte(sessionSecret))
	goth.UseProviders(
		google.New(googleKey, googleSecret, googleCallback),
	)

	auth := services.NewAuthenticator(userStore, jwtSigningKey)
	shortener := services.NewShortener(urlStore)
	renderer, err := services.NewRenderer()
	if err != nil {
		return err
	}

	router.InitRoutes(app, renderer, auth, shortener)

	return app.Start(fmt.Sprintf(":%d", port))
}
