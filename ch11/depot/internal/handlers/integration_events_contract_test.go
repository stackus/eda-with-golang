package handlers

import (
	"testing"

	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"

	"eda-in-golang/stores/storespb"
)

type String = matchers.String

var Like = matchers.Like

type Map = matchers.Map

func TestStoresConsumer(t *testing.T) {
	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "depot-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given   string
		expects string
		event   Map
	}{
		"AddStore": {
			given:   "a new store is created",
			expects: "a StoreCreated message",
			event: Map{
				"Name": String(storespb.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("NewStore"),
				}),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.NoError(t, pact.AddAsynchronousMessage().
				Given(tc.given).
				ExpectsToReceive(tc.expects).WithJSONContent(tc.event).
				ConsumedBy(func(contents v4.MessageContents) error {
					return nil
				}).
				Verify(t))
		})
	}
}
