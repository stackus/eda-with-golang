package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/notifications/internal/application"
	"github.com/stackus/eda-with-golang/ch4/notifications/notificationspb"
)

type server struct {
	app application.App
	notificationspb.UnimplementedNotificationsServiceServer
}

var _ notificationspb.NotificationsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	notificationspb.RegisterNotificationsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) NotifyOrderCreated(ctx context.Context, request *notificationspb.NotifyOrderCreatedRequest) (*notificationspb.NotifyOrderCreatedResponse, error) {
	err := s.app.NotifyOrderCreated(ctx, application.OrderCreated{
		SMSNumber: request.GetSmsNumber(),
		OrderID:   request.GetOrderId(),
	})
	return &notificationspb.NotifyOrderCreatedResponse{}, err
}
