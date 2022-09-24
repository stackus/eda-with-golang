package am

import (
	"context"
	"fmt"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"eda-in-golang/internal/ddd"
)

var tracer trace.Tracer
var propagator propagation.TextMapPropagator

func init() {
	tracer = otel.Tracer("internal/am")
	propagator = otel.GetTextMapPropagator()
}

func OtelMessageContextInjector() MessagePublisherMiddleware {
	return func(next MessagePublisher) MessagePublisher {
		return MessagePublisherFunc(func(ctx context.Context, topicName string, msg Message) error {
			var span trace.Span
			ctx, span = tracer.Start(ctx, msg.MessageName(), trace.WithSpanKind(trace.SpanKindProducer))
			propagator.Inject(ctx, MetadataCarrier(msg.Metadata()))
			defer span.End()

			err := next.Publish(ctx, topicName, msg)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			return err
		})
	}
}

func OtelMessageContextExtractor() MessageHandlerMiddleware {
	return func(next MessageHandler) MessageHandler {
		return MessageHandlerFunc(func(ctx context.Context, msg IncomingMessage) error {
			var span trace.Span
			ctx = propagator.Extract(ctx, MetadataCarrier(msg.Metadata()))
			ctx, span = tracer.Start(ctx, msg.MessageName(), trace.WithSpanKind(trace.SpanKindConsumer))
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

type MetadataCarrier ddd.Metadata

func (mc MetadataCarrier) Get(key string) string {
	switch v := ddd.Metadata(mc).Get(key).(type) {
	case nil:
		return ""
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (mc MetadataCarrier) Set(key, value string) {
	ddd.Metadata(mc).Set(key, value)
}

func (mc MetadataCarrier) Keys() []string {
	return ddd.Metadata(mc).Keys()
}
