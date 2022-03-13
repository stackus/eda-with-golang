package domain

type StoreCreated struct {
	Name     string
	Location string
}

func (StoreCreated) EventName() string { return "stores.StoreCreated" }

type StoreParticipationEnabled struct {
	Store *Store
}

func (StoreParticipationEnabled) EventName() string { return "stores.StoreParticipationEnabled" }

type StoreParticipationDisabled struct {
	Store *Store
}

func (StoreParticipationDisabled) EventName() string { return "stores.StoreParticipationDisabled" }

type StoreRebranded struct {
	Name string
}

func (StoreRebranded) EventName() string { return "stores.StoreRebranded" }
