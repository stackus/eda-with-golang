package jetstream

import (
	"context"
	"sync"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"eda-in-golang/ch7/internal/am"
)

type rawMessage struct {
	id       string
	name     string
	data     []byte
	acked    bool
	ackFn    func() error
	nackFn   func() error
	extendFn func() error
	killFn   func() error
}

type Stream struct {
	streamName string
	js         nats.JetStreamContext
	mu         sync.Mutex
}

var _ am.MessageStream[am.RawMessage, am.RawMessage] = (*Stream)(nil)
var _ am.RawMessage = (*rawMessage)(nil)

func NewStream(streamName string, js nats.JetStreamContext) *Stream {
	return &Stream{
		streamName: streamName,
		js:         js,
	}
}

func (s *Stream) Publish(ctx context.Context, topicName string, rawMsg am.RawMessage) error {
	data, err := proto.Marshal(&StreamMessage{
		Id:   rawMsg.ID(),
		Name: rawMsg.MessageName(),
		Data: rawMsg.Data(),
	})
	if err != nil {
		return err
	}

	_, err = s.js.PublishMsgAsync(&nats.Msg{
		Subject: topicName,
		Data:    data,
	})

	return err
}

func (s *Stream) Subscribe(topicName string, handler am.MessageHandler[am.RawMessage], options ...am.SubscriberOption) error {
	var err error

	s.mu.Lock()
	defer s.mu.Unlock()

	subCfg := am.NewSubscriberConfig(options)

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

	if ackType := subCfg.AckType(); ackType != am.AckTypeAuto {
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

func (s *Stream) handleMsg(cfg am.SubscriberConfig, handler am.MessageHandler[am.RawMessage]) func(*nats.Msg) {
	return func(natsMsg *nats.Msg) {

		var err error

		m := &StreamMessage{}
		err = proto.Unmarshal(natsMsg.Data, m)
		if err != nil {
			// TODO Nak? ... logging?
			return
		}

		msg := &rawMessage{
			id:       m.GetId(),
			name:     m.GetName(),
			data:     m.GetData(),
			acked:    false,
			ackFn:    func() error { return natsMsg.Ack() },
			nackFn:   func() error { return natsMsg.Nak() },
			extendFn: func() error { return natsMsg.InProgress() },
			killFn:   func() error { return natsMsg.Term() },
		}

		wCtx, cancel := context.WithTimeout(context.Background(), cfg.AckWait())
		defer cancel()

		errc := make(chan error)
		go func() {
			errc <- handler.HandleMessage(wCtx, msg)
		}()

		if cfg.AckType() == am.AckTypeAuto {
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

func (m rawMessage) ID() string {
	return m.id
}

func (m rawMessage) MessageName() string {
	return m.name
}

func (m rawMessage) Data() []byte {
	return m.data
}

func (m *rawMessage) Ack() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.ackFn()
}

func (m *rawMessage) NAck() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.nackFn()
}

func (m rawMessage) Extend() error {
	return m.extendFn()
}

func (m *rawMessage) Kill() error {
	if m.acked {
		return nil
	}

	m.acked = true
	return m.killFn()
}
