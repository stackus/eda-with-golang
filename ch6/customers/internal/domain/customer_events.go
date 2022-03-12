package domain

const (
	CustomerRegisteredEvent = "customers.CustomerRegistered"
	CustomerAuthorizedEvent = "customers.CustomerAuthorized"
	CustomerEnabledEvent    = "customers.CustomerEnabled"
	CustomerDisabledEvent   = "customers.CustomerDisabled"
)

type CustomerRegistered struct {
	Customer *Customer
}

type CustomerAuthorized struct {
	Customer *Customer
}

type CustomerEnabled struct {
	Customer *Customer
}

type CustomerDisabled struct {
	Customer *Customer
}
