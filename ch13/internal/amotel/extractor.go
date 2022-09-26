package amotel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"eda-in-golang/internal/am"
)

func OtelMessageContextExtractor() am.MessageHandlerMiddleware {
	return func(next am.MessageHandler) am.MessageHandler {
		return am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
			var span trace.Span
			ctx = propagator.Extract(ctx, MetadataCarrier(msg.Metadata()))
			ctx, span = tracer.Start(ctx, fmt.Sprintf("Receive(%s)", msg.MessageName()), trace.WithSpanKind(trace.SpanKindConsumer))
			defer span.End()

			err := next.HandleMessage(ctx, msg)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			return err
		})
	}
}
