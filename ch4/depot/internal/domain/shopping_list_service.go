package domain

import (
	"context"
)

type ShoppingListService struct {
	storeRepo   StoreRepository
	productRepo ProductRepository
}

func NewShoppingListService(storeRepo StoreRepository, productRepo ProductRepository) ShoppingListService {
	return ShoppingListService{
		storeRepo:   storeRepo,
		productRepo: productRepo,
	}
}

func (s ShoppingListService) BuildShoppingList(ctx context.Context, id, orderID string, items []OrderItem) (*ShoppingList, error) {
	stops := make([]*Stop, 0)
	storeIndexes := make(map[string]int)

	for _, item := range items {
		index, exists := storeIndexes[item.StoreID]
		if !exists {
			store, err := s.storeRepo.Find(ctx, item.StoreID)
			if err != nil {
				return nil, err
			}
			stops = append(stops, &Stop{
				StoreID:       item.StoreID,
				StoreName:     store.Name,
				StoreLocation: store.Location,
				Items:         make([]*Item, 0),
			})
			index = len(stops) - 1
		}

		product, err := s.productRepo.Find(ctx, item.StoreID, item.ProductID)
		if err != nil {
			return nil, err
		}
		stops[index].Items = append(stops[index].Items, &Item{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: item.Quantity,
		})
	}

	return &ShoppingList{
		ID:      id,
		OrderID: orderID,
		Stops:   stops,
		Status:  ShoppingListAvailable,
	}, nil
}
