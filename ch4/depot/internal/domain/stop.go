package domain

type Stops map[string]*Stop

type Stop struct {
	StoreName     string
	StoreLocation string
	Items         *Items
}

func (s *Stops) AddItem(store *Store, product *Product, quantity int) error {
	if _, exists := (*s)[store.ID]; !exists {
		(*s)[store.ID] = &Stop{
			StoreName:     store.Name,
			StoreLocation: store.Location,
			Items:         &Items{},
		}
	}

	return (*s)[store.ID].Items.AddItem(product, quantity)
}
