package tracing

import (
	"context"
	"fmt"

	tracing_config "github.com/ydb-platform/nbs/cloud/tasks/tracing/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
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

	fmt.Println("InitOpentelemetryTracing starting")
	fmt.Println("InitOpentelemetryTracing Config: %v", config)

	// TODO:_ what if tracing disabled?

	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:7881"))
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	// traceExporter, err := NewGRPCTraceExporter(ctx)
	if err != nil {
		fmt.Println("InitOpentelemetryTracing failed to create exporter")
		return nil, err
	}

	fmt.Println("InitOpentelemetryTracing created exporter")

	// XXXXXX serviceName := *config.ServiceName
	// XXXXXX fmt.Printf("InitOpentelemetryTracing config.ServiceName: %v\n", serviceName)
	// resource := otel_resource.NewWithAttributes(
	// 	semconv.SchemaURL,
	// 	semconv.ServiceNameKey.String("aaaaaa_service_name"),
	// 	// TODO:_ hostname?
	// )

	fmt.Println("InitOpentelemetryTracing created resource")

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter), // TODO:_ timeout?
		// sdktrace.WithResource(resource),
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

	return tracerProvider.Shutdown, nil
}
