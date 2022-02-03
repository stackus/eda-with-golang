package domain

type Stop struct {
	StoreID       string
	StoreName     string
	StoreLocation string
	Items         []*Item
}
