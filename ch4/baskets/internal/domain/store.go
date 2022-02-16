package domain

type StoreID string

type Store struct {
	ID       StoreID
	Name     string
	Location string
}

func (i StoreID) String() string {
	return string(i)
}
