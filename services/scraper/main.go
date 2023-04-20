package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/pkg/database"
	"github.com/kanade0404/tenhou-log/services/scraper/api"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
)

func main() {
	log.Println("Run start")
	ctx := context.Background()
	env, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	exporter, err := texporter.New(texporter.WithProjectID(""))
	if err != nil {
		log.Fatalln(err)
	}
	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(os.Getenv("K_SERVICE")),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	defer func(tp *sdktrace.TracerProvider, ctx context.Context) {
		err := tp.ForceFlush(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}(tp, ctx)
	otel.SetTracerProvider(tp)

	tracer := otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper")
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()
	client, err := database.CreateClient(ctx, env.Dialect(), env.ConnectionString())
	if err != nil {
		log.Fatalf("failed database connection: %v", err)
	}
	if err := database.Migration(ctx, client); err != nil {
		log.Fatalln(err)
	}
	router := chi.NewRouter()
	a := api.New(ctx, client, env)
	router.Post("/scraping", a.Scraper)
	port := env.Port()
	if port != "" {
		port = "8088"
	}
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}
	log.Printf("server listen %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed listening server: %v", err)
	}
}
