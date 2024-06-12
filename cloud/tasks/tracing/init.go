package tracing

import (
	"context"

	tracing_config "github.com/ydb-platform/nbs/cloud/tasks/tracing/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	otel_resource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type tracerFieldKey struct{}

const (
	tracerName = "yc-disk-manager"
)

// TODO:_ put it in context?

func GetTracer(ctx context.Context) trace.Tracer {
	return ctx.Value(tracerFieldKey{}).(trace.Tracer)
}

func InitOpentelemetryTracing(
	ctx context.Context,
	config *tracing_config.TracingConfig,
) (context.Context, error) {

	// TODO:_ what if tracing disabled?

	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return ctx, err
	}

	// traceExporter, err := otlptracegrpc.New(
	// 	ctx,
	// 	otlptracegrpc.WithEndpoint("localhost:7881"),
	// 	otlptracegrpc.WithInsecure(),
	// )
	// if err != nil {
	// 	return err
	// }

	resource := otel_resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(*config.ServiceName),
		// TODO:_ hostname?
	)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter), // TODO:_ timeout?
		sdktrace.WithResource(resource),
	)
	otel.SetTracerProvider(tracerProvider)

	ctx = context.WithValue(ctx, tracerFieldKey{}, otel.Tracer(tracerName))

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return ctx, err
}
