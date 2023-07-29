package tracer

import (
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/logger"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/kanade0404/tenhou-log/pkg/config"
)

func InitTracer(ctx context.Context, env *config.Config) error {
	logger.Info("InitTracer start")
	defer logger.Info("InitTracer end")
	exporter, err := texporter.New(texporter.WithProjectID(env.ProjectID()))
	if err != nil {
		return err
	}
	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(os.Getenv("K_SERVICE")),
		),
	)
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	defer func(tp *sdktrace.TracerProvider, ctx context.Context) {
		err := tp.ForceFlush(ctx)
		if err != nil {
			logger.ErrorF("failed to *trace.TracerProvider.ForceFlush(). err: %v", err.Error())
		}
	}(tp, ctx)
	otel.SetTracerProvider(tp)
	return nil
}
func NewTracer(name string) trace.Tracer {
	return otel.GetTracerProvider().Tracer(fmt.Sprintf("github.com/kanade0404/tenhou-log/%s", name))
}
