package handlers

import (
	"encoding/json"
	"testing"

	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"

	"eda-in-golang/internal/registry"
	"eda-in-golang/internal/registry/serdes"
	"eda-in-golang/stores/storespb"
)

func TestStoresProducer(t *testing.T) {
	var err error

	type rawMessage struct {
		Name    string
		Payload json.RawMessage
	}

	reg := registry.New()
	err = storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	verifier := message.Verifier{}
	err = verifier.Verify(t, message.VerifyMessageRequest{
		VerifyRequest: provider.VerifyRequest{
			Provider:                   "stores-pub",
			ProviderVersion:            "1.0.0",
			BrokerURL:                  "http://127.0.0.1:9292",
			BrokerUsername:             "pactuser",
			BrokerPassword:             "pactpass",
			PublishVerificationResults: true,
			FailIfNoPactsFound:         true,
		},
		MessageHandlers: map[string]message.Handler{
			"a StoreCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				msg := rawMessage{
					Name: storespb.StoreCreatedEvent,
					Payload: reg.MustSerialize(storespb.StoreCreatedEvent, &storespb.StoreCreated{
						Id:       "store-id",
						Name:     "NewStore",
						Location: "NewLocation",
					}),
				}

				return msg, nil, nil
			},
			"a StoreRebranded message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				msg := rawMessage{
					Name: storespb.StoreRebrandedEvent,
					Payload: reg.MustSerialize(storespb.StoreRebrandedEvent, &storespb.StoreRebranded{
						Id:   "store-id",
						Name: "RebrandedStore",
					}),
				}

				return msg, nil, nil
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
