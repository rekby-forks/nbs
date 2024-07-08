package tracing

import (
	"context"

	"github.com/ydb-platform/nbs/cloud/tasks/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	grpc_metadata "google.golang.org/grpc/metadata"
)

func ExtractTraceContext(
	ctx context.Context,
) context.Context {

	logging.Info(ctx, "ExtractTraceContext starting")

	// TODO:_ should we depend on grpc_metadata here? Maybe we should call smth from headers package here.
	md, ok := grpc_metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	logging.Info(ctx, "ExtractTraceContext got grpc metadata: %v", md)

	// TODO:_ should we handle absent traceparent somehow?
	// TODO:_ put fields names into constants
	vals := md.Get("traceparent")
	if len(vals) == 0 {
		logging.Info(ctx, "ExtractTraceContext no traceparent, finishing")
		return ctx
	}
	logging.Info(ctx, "ExtractTraceContext got traceparent: %v", vals)

	headers := make(map[string]string)
	headers["traceparent"] = vals[0]

	// TODO:_ comment about traceparent/tracestate contract + link do docs?
	vals = md.Get("tracestate")
	if len(vals) != 0 && len(vals[0]) != 0 {
		logging.Info(ctx, "ExtractTraceContext got tracestate: %v", vals)
		headers["tracestate"] = vals[0]
	}

	logging.Info(ctx, "ExtractTraceContext extracting from headers to ctx, headers: %v", headers)
	prop := otel.GetTextMapPropagator()
	ctx = prop.Extract(ctx, propagation.MapCarrier(headers))
	logging.Info(ctx, "ExtractTraceContext extracted from headers to ctx, finishing")
	return ctx
}

func InjectTraceContext(ctx context.Context) context.Context {
	mapCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, mapCarrier)

	// TODO:_ here also: do not use grpc metadata directly, call something from headers package. But not Append.
	// TODO:_ why not outgoing?
	existingMd, ok := grpc_metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	traceparent, ok := mapCarrier["traceparent"]
	if !ok {
		// TODO:_ can it happen !ok?
		logging.Info(ctx, "InjectTraceContext no traceparent, WTF, finishing")
		return ctx
	}
	existingMd.Set("traceparent", traceparent)
	logging.Info(ctx, "InjectTraceContext injected traceparent: %v", traceparent)

	tracestate, ok := mapCarrier["tracestate"]
	if ok {
		existingMd.Set("tracestate", tracestate)
		logging.Info(ctx, "InjectTraceContext injected tracestate: %v", tracestate)
	} else {
		existingMd.Delete("tracestate")
		logging.Info(ctx, "InjectTraceContext: no tracestate, deleted tracestate from metadata")
	}

	ctx = grpc_metadata.NewIncomingContext(ctx, existingMd)
	return grpc_metadata.NewOutgoingContext(ctx, existingMd)
}
