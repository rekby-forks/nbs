package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	grpc_metadata "google.golang.org/grpc/metadata"
)

func ExtractTraceContext(
	ctx context.Context,
) context.Context {

	// TODO:_ should we depend on grpc_metadata here? Maybe we should call smth from headers package here.
	md, ok := grpc_metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	// TODO:_ should we handle absent traceparent somehow?
	// TODO:_ put fields names into constants
	vals := md.Get("traceparent")
	if len(vals) == 0 {
		return ctx
	}

	headers := make(map[string]string)
	headers["traceparent"] = vals[0]

	// TODO:_ comment about traceparent/tracestate contract + link do docs?
	vals = md.Get("tracestate")
	if len(vals) != 0 && len(vals[0]) != 0 {
		headers["tracestate"] = vals[0]
	}

	prop := otel.GetTextMapPropagator()
	ctx = prop.Extract(ctx, propagation.MapCarrier(headers))
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
		return ctx
	}
	existingMd.Set("traceparent", traceparent)

	tracestate, ok := mapCarrier["tracestate"]
	if ok {
		existingMd.Set("tracestate", tracestate)
	} else {
		existingMd.Delete("traceparent")
	}

	return grpc_metadata.NewOutgoingContext(ctx, existingMd)
}
