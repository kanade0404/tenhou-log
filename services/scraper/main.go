package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kanade0404/tenhou-log/pkg/config"
	tracer2 "github.com/kanade0404/tenhou-log/pkg/driver/tracer"
	"github.com/kanade0404/tenhou-log/pkg/logger"
	"github.com/kanade0404/tenhou-log/services/scraper/api"
)

func main() {
	logger.Init()
	logger.Info("Run start")
	ctx := context.Background()
	env, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	err = tracer2.InitTracer(ctx, env)
	if err != nil {
		log.Fatalln(err)
	}
	tracer := tracer2.NewTracer("services/scraper")
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()
	router := chi.NewRouter()
	a := api.New(ctx, env)
	router.Post("/scraping", a.Scraper)
	port := env.Port()
	if port != "" {
		port = "8088"
	}
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}
	logger.InfoF("server listen %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.FatalF("failed listening server: %v", err)
	}
}
