package api

import (
	"context"

	"go.opentelemetry.io/otel"

	"github.com/kanade0404/tenhou-log/pkg/config"
)

var apiTracer = otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper/api/scraper")

type Api struct {
	context.Context
	config *config.Config
}

func New(c context.Context, cnf *config.Config) *Api {
	ctx, span := apiTracer.Start(c, "New")
	defer span.End()
	return &Api{
		ctx,
		cnf,
	}
}
