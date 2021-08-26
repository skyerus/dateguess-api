package main

import (
	"dateguess-api/internal/app"
	"dateguess-api/internal/handler"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := app.Init()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}
	defer func(logger *zap.SugaredLogger) {
		err := logger.Sync()
		if err != nil {
			log.Fatalf("failed to sync logger: %v", err)
		}
	}(a.Logger)

	port := "8080"

	a.Logger.Infof("dateguess-api - Listening on port %s:", port)

	a.Logger.Fatal(
		http.ListenAndServe(
			":"+port,
			routes(a),
		).Error(),
	)
}

func routes(app *app.App) http.Handler {
	r := mux.NewRouter()
	r.Use(app.Middleware.Cors)

	eventHandler := handler.NewEvent(app.Router, app.Logger, app.EventService)
	articleHandler := handler.NewArticle(app.Router, app.Logger, app.ArticleService)

	r.HandleFunc("/", handler.NewHealth(app.Router).Get).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/random_article", articleHandler.Random).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/random_historical_event", eventHandler.RandomHistoricalEvent).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/random_historical_events", eventHandler.RandomHistoricalEvents).Methods("GET", "OPTIONS")

	return r
}
