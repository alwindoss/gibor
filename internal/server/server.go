package server

import (
	"io/fs"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/alwindoss/gibor/ui"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Config struct {
}

func Run(cfg *Config) error {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Register backend APIs first.
	api := e.Group("/api")
	api.GET("/health", func(c *echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	// Register embedded frontend.
	frontend, err := ui.FS()
	if err != nil {
		log.Fatal(err)
	}

	registerFrontend(e, frontend)

	// Start server.
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}

func registerFrontend(e *echo.Echo, frontend fs.FS) {
	serve := func(c *echo.Context) error {
		requestPath := strings.TrimPrefix(c.Request().URL.Path, "/")

		if requestPath == "" {
			requestPath = "index.html"
		}

		// Never serve the SPA for API requests.
		if strings.HasPrefix(requestPath, "api/") ||
			requestPath == "api" {
			return echo.ErrNotFound
		}

		cleanPath := path.Clean(requestPath)

		// Prevent invalid filesystem paths.
		if !fs.ValidPath(cleanPath) {
			return echo.ErrNotFound
		}

		data, err := fs.ReadFile(frontend, cleanPath)

		if err != nil {
			// Missing files with extensions are actual
			// missing resources, not React routes.
			if path.Ext(cleanPath) != "" {
				return echo.ErrNotFound
			}

			// React Router fallback.
			data, err = fs.ReadFile(frontend, "index.html")
			if err != nil {
				return echo.ErrNotFound
			}
			cleanPath = "index.html"
		}

		contentType := mime.TypeByExtension(
			path.Ext(cleanPath),
		)

		if contentType == "" {
			contentType = "application/octet-stream"
		}

		return c.Blob(
			http.StatusOK,
			contentType,
			data,
		)
	}

	e.GET("/", serve)
	e.GET("/*", serve)
}
