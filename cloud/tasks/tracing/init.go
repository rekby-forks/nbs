package tracing

import (
	"context"
	"fmt"
	"os"
	"time"

	tracing_config "github.com/ydb-platform/nbs/cloud/tasks/tracing/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
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

// func NewStdoutTraceExporter() (*otlptrace.Exporter, error) {
// 	return stdouttrace.New(stdouttrace.WithPrettyPrint())
// }

// func NewGRPCTraceExporter(ctx context.Context) (*otlptrace.Exporter, error) {
// 	return otlptracegrpc.New(
// 		ctx,
// 		otlptracegrpc.WithEndpoint("localhost:7881"),
// 		otlptracegrpc.WithInsecure(),
// 	)
// }

// func NewHTTPTraceExporter() (*otlptrace.Exporter, error) {
// }

func InitOpentelemetryTracing(
	ctx context.Context,
	config *tracing_config.TracingConfig,
) (shutdown func(context.Context) error, err error) {

	fmt.Println("CHECK InitOpentelemetryTracing starting")
	fmt.Printf("CHECK InitOpentelemetryTracing Config: %v\n", config)
	goLogPings(ctx)

	// TODO:_ what if tracing disabled?

	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:7881"))
	// traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	// traceExporter, err := NewGRPCTraceExporter(ctx)
	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("localhost:7881"),
		otlptracegrpc.WithInsecure(),
	)

	if err != nil {
		fmt.Println("CHECK InitOpentelemetryTracing failed to create exporter")
		return nil, err
	}

	fmt.Println("CHECK InitOpentelemetryTracing created exporter")

	// TODO:_ write this normally
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}

	serviceName := *config.ServiceName
	fmt.Printf("InitOpentelemetryTracing config.ServiceName: %v\n", serviceName)
	resource := otel_resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.HostNameKey.String(hostname), // TODO:_ hostname?
		attribute.String("generation", "CCC"),
	)

	fmt.Println("CHECK InitOpentelemetryTracing created resource")

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter), // TODO:_ timeout?
		sdktrace.WithResource(resource),
	)
	fmt.Println("CHECK InitOpentelemetryTracing created trace provider")
	otel.SetTracerProvider(tracerProvider)
	fmt.Println("CHECK InitOpentelemetryTracing set trace provider")

	// ctx = context.WithValue(ctx, tracerFieldKey{}, otel.Tracer(tracerName))
	// TODO:_ should I create separate tracer for each operation?
	tracer = otel.Tracer(tracerName)
	fmt.Println("CHECK InitOpentelemetryTracing created tracer")

	otel.SetTextMapPropagator(propagation.TraceContext{})
	fmt.Println("CHECK InitOpentelemetryTracing set propagator, finishing")

	return tracerProvider.Shutdown, nil
}

func goLogPings(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Millisecond)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("CHECK Context cancelled, finishing trace pings")
			return
		case <-ticker.C:
			fmt.Println("CHECK PING from trace pings")
		}
	}()
}
