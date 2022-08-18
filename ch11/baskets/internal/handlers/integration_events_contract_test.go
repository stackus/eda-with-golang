package handlers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/registry/serdes"
	"eda-in-golang/stores/storespb"
)

type String = matchers.String
type Map = matchers.Map

var Like = matchers.Like

func TestStoresConsumer(t *testing.T) {
	type mocks struct {
		stores   *domain.MockStoreCacheRepository
		products *domain.MockProductCacheRepository
	}

	reg := registry.New()
	err := storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "baskets-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given   []models.ProviderState
		expects string
		message Map
		on      func(m mocks)
	}{
		"AddStore": {
			expects: "a StoreCreated message",
			message: Map{
				"Name": String(storespb.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("NewStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore").Return(nil)
			},
		},
		"RebrandStore": {
			expects: "a StoreRebranded message",
			message: Map{
				"Name": String(storespb.StoreRebrandedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("RebrandedStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Rename", mock.Anything, "store-id", "RebrandedStore").Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   domain.NewMockStoreCacheRepository(t),
				products: domain.NewMockProductCacheRepository(t),
			}
			if tc.on != nil {
				tc.on(m)
			}
			handlers := NewIntegrationEventHandlers(m.stores, m.products)

			message := pact.AddAsynchronousMessage()
			for _, given := range tc.given {
				message = message.GivenWithParameter(given)
			}
			assert.NoError(t, message.
				ExpectsToReceive(tc.expects).
				WithJSONContent(tc.message).
				ConsumedBy(func(contents v4.MessageContents) error {
					message := contents.Content.(map[string]any)

					data, err := json.Marshal(message["Payload"])
					if err != nil {
						return err
					}
					payload, err := reg.Deserialize(message["Name"].(string), data)
					if err != nil {
						return err
					}

					return handlers.HandleEvent(context.Background(), ddd.NewEvent(message["Name"].(string), payload))
				}).
				Verify(t))
		})
	}
}
