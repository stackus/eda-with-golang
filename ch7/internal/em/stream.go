package em

import (
	"eda-in-golang/ch7/internal/ddd"
)

type Stream interface {
	Publish(topicName string, event ddd.Event, options ...PublisherOption) error
	Subscribe(topicName string, handler MessageHandler, options ...SubscriberOption) error
}
