package amprom

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"eda-in-golang/internal/am"
)

func ReceivedMessagesCounter(serviceName string) am.MessageHandlerMiddleware {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: serviceName,
		Name:      "received_messages",
		Help:      fmt.Sprintf("The total number of messages received by %s", serviceName),
	})
	return func(next am.MessageHandler) am.MessageHandler {
		return am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
			counter.Inc()
			return next.HandleMessage(ctx, msg)
		})
	}
}
