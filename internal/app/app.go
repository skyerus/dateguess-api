package app

import (
	articleRepo "dateguess-api/internal/repository/article"
	eventRepo "dateguess-api/internal/repository/event"
	"dateguess-api/internal/service/article"
	"dateguess-api/internal/service/event"
	"dateguess-api/pkg/database"
	"dateguess-api/pkg/logger"
	"dateguess-api/pkg/middleware"
	"dateguess-api/pkg/router"
	"fmt"
	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

type App struct {
	// Env holds the environment variables passed to the app
	Env Environment
	// Router handles responding to http requests
	Router *router.Router
	// Logger is the logger for the service
	Logger *zap.SugaredLogger
	// Middleware is the middleware of the app
	Middleware *middleware.Middleware
	// EventService provides an API for interacting with events
	EventService *event.Service
	// EventService provides an API for interacting with events data store
	EventRepository *eventRepo.Repository
	// ArticleService provides an API for interacting with articles
	ArticleService *article.Service
}

func Init() (*App, error) {
	var app App
	var err error

	if err = env.Parse(&app.Env); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	l, err := logger.New(app.Env.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}
	app.Logger = l.Sugar()

	app.Router = router.New(
		router.Config{
			DefaultHeaders: map[string][]string{
				"Content-Type": {"application/json"},
			},
			RequestIDContextKey: "requestID",
		},
		app.Logger,
	)

	db, err := database.New(
		database.Config{
			MySQLUser:     app.Env.MySQLUser,
			MySQLPassword: app.Env.MySQLPassword,
			MySQLHost:     app.Env.MySQLHost,
			MySQLOptions:  "parseTime=1&sql_mode=TRADITIONAL",
			MySQLTimeZone: "\"+0:00\"",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	app.EventRepository = eventRepo.NewRepository(db, app.Logger)
	app.EventService = event.NewService(app.EventRepository)

	app.Middleware = middleware.NewMiddleware(app.Env.AllowOrigin, app.Env.Env == "dev")

	aRepo := articleRepo.NewRepository(
		app.Env.GuardianKey,
		app.Env.GuardianBaseURL,
		app.Logger,
	)
	app.ArticleService = article.NewService(aRepo)

	return &app, nil
}
