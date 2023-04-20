package api

import (
	"context"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/services/ent"
	"go.opentelemetry.io/otel"
)

var apiTracer = otel.GetTracerProvider().Tracer("github.com/kanade0404/tenhou-log/services/scraper/api/scraper")

type Api struct {
	context.Context
	client *ent.Client
	config *config.Config
}

func New(c context.Context, client *ent.Client, cnf *config.Config) *Api {
	ctx, span := apiTracer.Start(c, "New")
	defer span.End()
	return &Api{
		ctx,
		client,
		cnf,
	}
}
