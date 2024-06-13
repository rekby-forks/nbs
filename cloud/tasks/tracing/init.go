package tracing

import (
	"context"
	"fmt"

	tracing_config "github.com/ydb-platform/nbs/cloud/tasks/tracing/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	otel_resource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// type tracerFieldKey struct{}

const (
	tracerName = "yc-disk-manager"
)

// func GetTracer(ctx context.Context) trace.Tracer {
// 	return ctx.Value(tracerFieldKey{}).(trace.Tracer)
// }

var tracer trace.Tracer

func GetTracer() trace.Tracer {
	return tracer
}

func InitOpentelemetryTracing(
	ctx context.Context,
	config *tracing_config.TracingConfig,
) error {

	fmt.Println("InitOpentelemetryTracing starting")

	// TODO:_ what if tracing disabled?

	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		fmt.Println("InitOpentelemetryTracing failed to create exporter")
		return err
	}

	fmt.Println("InitOpentelemetryTracing created exporter")

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

	fmt.Println("InitOpentelemetryTracing created resource")

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter), // TODO:_ timeout?
		sdktrace.WithResource(resource),
	)
	fmt.Println("InitOpentelemetryTracing created trace provider")
	otel.SetTracerProvider(tracerProvider)
	fmt.Println("InitOpentelemetryTracing set trace provider")

	// ctx = context.WithValue(ctx, tracerFieldKey{}, otel.Tracer(tracerName))
	// TODO:_ should I create separate tracer for each operation?
	tracer = otel.Tracer(tracerName)
	fmt.Println("InitOpentelemetryTracing created tracer")

	otel.SetTextMapPropagator(propagation.TraceContext{})
	fmt.Println("InitOpentelemetryTracing set propagator, finishing")

	return nil
}
