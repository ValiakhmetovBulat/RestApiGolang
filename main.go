package main

import (
	_ "RestApiGolang/docs"
	"RestApiGolang/internal/config"
	"RestApiGolang/internal/database/sqlite"
	"RestApiGolang/internal/http-server/api"
	"RestApiGolang/internal/http-server/router"
	log "RestApiGolang/internal/logger"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

const (
	configPath = "config/local.yaml"
)

//	@title			Rest API Golang
//	@version		1.0
//	@description	Rest API Server Golang

//	@host		localhost:8888
//	@basePath	/

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

	var routes = []router.Route{
		{
			Method:      router.GET,
			Path:        "/",
			HandlerFunc: api.RootHandler,
		},
		{
			Method:      router.GET,
			Path:        "/urls",
			HandlerFunc: api.GetUrls,
		},
		{
			Method:      router.GET,
			Path:        "/url/{alias}",
			HandlerFunc: api.GetUrlByAlias,
		},
		{
			Method:      router.POST,
			Path:        "/urls",
			HandlerFunc: api.PostUrl,
		},
		{
			Method:      router.PUT,
			Path:        "/url",
			HandlerFunc: api.PutUrl,
		},
		{
			Method:      router.DELETE,
			Path:        "/url",
			HandlerFunc: api.DeleteUrlByAlias,
		},
		{
			Method:      router.GET,
			Path:        "/{alias}",
			HandlerFunc: api.Redirect,
		},
	}

	r := router.NewRouter(routes)

	r.Mount("/swagger", httpSwagger.WrapHandler)

	log.WithFields(logrus.Fields{
		"address": cfg.Address,
	}).Info("starting server")

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}

	log.Error("server stopped")
}
