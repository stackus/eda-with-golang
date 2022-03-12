package domain

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
)

type StoreCreated struct {
	Store *Store
}

type StoreParticipationEnabled struct {
	Store *Store
}

type StoreParticipationDisabled struct {
	Store *Store
}
