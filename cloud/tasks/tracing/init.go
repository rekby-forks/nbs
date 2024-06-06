package tracing

import (
	"context"

	tracing_config "github.com/ydb-platform/nbs/cloud/tasks/tracing/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	otel_resource "go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	tracerName = "yc-disk-manager"
)

func getTracer(ctx context.Context) {
	otel.Tracer(tracerName)
}

func InitOpentelemetryTracing(
	ctx context.Context,
	config *tracing_config.TracingConfig,
) error {

	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("localhost:7881"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	resource := otel_resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("gayurgin_debug"),
		// TODO:_ hostname?
	)

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter), // TODO:_ timeout?
		trace.WithResource(resource),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return nil
}
