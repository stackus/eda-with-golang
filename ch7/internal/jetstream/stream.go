package jetstream

import (
	"context"
	"sync"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"eda-in-golang/ch7/internal/ddd"
	"eda-in-golang/ch7/internal/em"
	"eda-in-golang/ch7/internal/registry"
)

type Stream struct {
	streamName string
	js         nats.JetStreamContext
	reg        registry.Registry
	mu         sync.Mutex
}

var _ em.Stream = (*Stream)(nil)

func NewStream(streamName string, js nats.JetStreamContext, reg registry.Registry) *Stream {
	return &Stream{
		streamName: streamName,
		js:         js,
		reg:        reg,
	}
}

func (s *Stream) Publish(topicName string, event ddd.Event, options ...em.PublisherOption) error {
	data, err := s.reg.Serialize(event.EventName(), event.Payload())
	if err != nil {
		return err
	}

	pubCfg := em.NewPublisherConfig(options...)

	headers := make(map[string]*EventMessage_Values)
	for key, values := range pubCfg.Headers() {
		headers[key] = &EventMessage_Values{Values: values}
	}

	var m []byte

	m, err = proto.Marshal(&EventMessage{
		Id:               event.ID(),
		EventName:        event.EventName(),
		AggregateName:    event.AggregateName(),
		AggregateId:      event.AggregateID(),
		AggregateVersion: int32(event.AggregateVersion()),
		OccurredAt:       timestamppb.New(event.OccurredAt()),
		Headers:          headers,
		Data:             data,
	})
	if err != nil {
		return err
	}

	_, err = s.js.PublishMsgAsync(&nats.Msg{
		Subject: topicName,
		Data:    m,
	})

	return err
}

func (s *Stream) Subscribe(topicName string, handler em.MessageHandler, options ...em.SubscriberOption) error {
	var err error

	s.mu.Lock()
	defer s.mu.Unlock()

	subCfg := em.NewSubscriberConfig(options...)

	opts := []nats.SubOpt{
		nats.MaxDeliver(subCfg.MaxRedeliver()),
	}
	cfg := &nats.ConsumerConfig{
		MaxDeliver: subCfg.MaxRedeliver(),
	}
	if groupName := subCfg.GroupName(); groupName != "" {
		cfg.DeliverSubject = groupName
		cfg.DeliverGroup = groupName
		cfg.Durable = groupName

		opts = append(opts, nats.Bind(s.streamName, groupName), nats.Durable(groupName))
	}

	if ackType := subCfg.AckType(); ackType != em.AckTypeAuto {
		ackWait := subCfg.AckWait()

		cfg.AckPolicy = nats.AckExplicitPolicy
		cfg.AckWait = ackWait

		opts = append(opts, nats.AckExplicit(), nats.AckWait(ackWait))
	} else {
		cfg.AckPolicy = nats.AckNonePolicy
		opts = append(opts, nats.AckNone())
	}

	_, err = s.js.AddConsumer(s.streamName, cfg)
	if err != nil {
		return err
	}

	if groupName := subCfg.GroupName(); groupName == "" {
		_, err = s.js.Subscribe(topicName, s.handleMsg(subCfg, handler), opts...)
	} else {
		_, err = s.js.QueueSubscribe(topicName, groupName, s.handleMsg(subCfg, handler), opts...)
	}

	return nil
}

func (s *Stream) handleMsg(cfg em.SubscriberConfig, handler em.MessageHandler) func(*nats.Msg) {
	return func(natsMsg *nats.Msg) {

		var err error

		m := &EventMessage{}
		err = proto.Unmarshal(natsMsg.Data, m)
		if err != nil {
			// TODO Nak? ... logging?
			return
		}

		headers := map[string][]string{}
		for key, values := range m.GetHeaders() {
			headers[key] = values.GetValues()
		}

		msg := &message{
			reg:        s.reg,
			eventID:    m.GetId(),
			eventName:  m.GetEventName(),
			data:       m.GetData(),
			payload:    nil,
			aggID:      m.GetAggregateId(),
			aggName:    m.GetAggregateName(),
			aggVersion: int(m.GetAggregateVersion()),
			occurredAt: m.GetOccurredAt().AsTime(),
			headers:    headers,
			acked:      false,
			ackFn:      func() error { return natsMsg.Ack() },
			nackFn:     func() error { return natsMsg.Nak() },
			extendFn:   func() error { return natsMsg.InProgress() },
			killFn:     func() error { return natsMsg.Term() },
		}

		wCtx, cancel := context.WithTimeout(context.Background(), cfg.AckWait())
		defer cancel()

		errc := make(chan error)
		go func() {
			errc <- handler.HandleMessage(wCtx, msg)
		}()

		if cfg.AckType() == em.AckTypeAuto {
			err = msg.Ack()
			if err != nil {
				// TODO logging?
			}
		}

		select {
		case err = <-errc:
			if err == nil {
				if ackErr := msg.Ack(); ackErr != nil {
					// TODO logging?
				}
				return
			}
			if nakErr := msg.NAck(); nakErr != nil {
				// TODO logging?
			}
		case <-wCtx.Done():
			// TODO logging?
			return
		}
	}
}

// func (s *Stream) Start(ctx context.Context) error {
// 	var err error
// 	s.js, err = s.nc.JetStream(nats.ContextOpt{Context: ctx})
// 	if err != nil {
// 		return err
// 	}
//
// 	if s.cfg.Name == "" {
// 		return fmt.Errorf("cannot start stream without a name")
// 	}
//
// 	info, err := s.js.StreamInfo(s.cfg.Name, nats.ContextOpt{Context: ctx})
// 	if err != nil {
// 		return err
// 	}
//
// 	// create the stream if it does not exist
// 	if info == nil {
// 		if len(s.cfg.Subjects) == 0 {
// 			s.cfg.Subjects = []string{fmt.Sprintf("%s.>", s.cfg.Name)}
// 		}
//
// 		_, err = s.js.AddStream(s.cfg, nats.ContextOpt{Context: ctx})
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	for _, subscribers := range s.subs {
// 		for _, sub := range subscribers {
// 			err = sub.subscribe(ctx, s.js)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
//
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		case <-s.shutdown:
// 			return nil
// 		default:
// 		}
// 	}
// }
//
// func (s *Stream) Shutdown() (err error) {
// 	close(s.shutdown)
// 	return s.nc.Drain()
// }
