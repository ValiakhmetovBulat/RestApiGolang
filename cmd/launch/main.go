package main

import (
	_ "RestApiGolang/cmd/launch/docs"
	"RestApiGolang/internal/config"
	"RestApiGolang/internal/database/sqlite"
	"RestApiGolang/internal/http-server/api"
	mwLogger "RestApiGolang/internal/http-server/middleware/logger"
	log "RestApiGolang/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

const (
	configPath = "../../config/local.yaml"
)

// @title Rest API Golang
// @version 1.0
// @description Rest API Server Golang

// @host localhost:8888
// @basePath /

func main() {
	// load config
	cfg := config.MustLoad(configPath)

	// setup logger
	log.SetupLogger(cfg.Env)

	// log app starting info
	log.WithFields(logrus.Fields{
		"env":     cfg.Env,
		"version": "v1.0",
	}).Info("starting app")

	log.Debug("debug mode enabled")

	err := sqlite.Setup(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log.Logger))
	// router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)

	router.Get("/", api.RootHandler)
	router.Get("/urls", api.GetUrls)
	//router.Get("/url", api.GetUrlByAlias)
	//router.Post("/urls", api.PostUrl)
	//router.Delete("/url", api.DeleteUrlByAlias)
	//router.Put("/url", api.PutUrl)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8888/swagger/swagger.json"),
	))

	log.WithFields(logrus.Fields{
		"address": cfg.Address,
	}).Info("starting server")

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}

	log.Error("server stopped")
}
