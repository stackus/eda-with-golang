package amprom

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"eda-in-golang/internal/am"
)

func SentMessagesCounter(serviceName string) am.MessagePublisherMiddleware {
	metric := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: serviceName,
		Name:      "sent_messages",
		Help:      fmt.Sprintf("The total number of messages sent by %s", serviceName),
	})
	return func(next am.MessagePublisher) am.MessagePublisher {
		return am.MessagePublisherFunc(func(ctx context.Context, topicName string, msg am.Message) error {
			metric.Inc()
			return next.Publish(ctx, topicName, msg)
		})
	}
}
