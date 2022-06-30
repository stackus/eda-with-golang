package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/es"
)

func TestApplication_AddItem(t *testing.T) {
	product := &domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	store := &domain.Store{
		ID:   "store-id",
		Name: "store-name",
	}

	type fields struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}
	type args struct {
		ctx context.Context
		add AddItem
	}
	tests := map[string]struct {
		args    args
		on      func(f fields)
		wantErr bool
	}{
		"AddItem": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(store, nil)
				f.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
			},
		},
		"AddItem.NoProduct": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(nil, fmt.Errorf("no product"))
			},
			wantErr: true,
		},
		"AddItem.NoStore": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(nil, fmt.Errorf("no store"))
			},
			wantErr: true,
		},
		"AddItem.SaveFailed": {
			args: args{
				ctx: context.Background(),
				add: AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(store, nil)
				f.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(fmt.Errorf("save failed"))
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f := fields{
				baskets:   domain.NewMockBasketRepository(t),
				stores:    domain.NewMockStoreRepository(t),
				products:  domain.NewMockProductRepository(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
			}
			a := Application{
				baskets:   f.baskets,
				stores:    f.stores,
				products:  f.products,
				publisher: f.publisher,
			}
			if tt.on != nil {
				tt.on(f)
			}

			if err := a.AddItem(tt.args.ctx, tt.args.add); (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApplication_CancelBasket(t *testing.T) {
	type fields struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}
	type args struct {
		ctx    context.Context
		cancel CancelBasket
	}
	tests := map[string]struct {
		args    args
		on      func(f fields)
		wantErr bool
	}{
		"CancelBasket": {
			args: args{
				ctx: context.Background(),
				cancel: CancelBasket{
					ID: "basket-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
				f.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
				f.publisher.On("Publish", context.Background(), mock.AnythingOfType("ddd.event")).Return(nil)
			},
		},
		"CancelBasket.CheckedOut": {
			args: args{
				ctx: context.Background(),
				cancel: CancelBasket{
					ID: "basket-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsCheckedOut,
				}, nil)
			},
			wantErr: true,
		},
		"CancelBasket.Canceled": {
			args: args{
				ctx: context.Background(),
				cancel: CancelBasket{
					ID: "basket-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsCanceled,
				}, nil)
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f := fields{
				baskets:   domain.NewMockBasketRepository(t),
				stores:    domain.NewMockStoreRepository(t),
				products:  domain.NewMockProductRepository(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
			}
			a := Application{
				baskets:   f.baskets,
				stores:    f.stores,
				products:  f.products,
				publisher: f.publisher,
			}
			if tt.on != nil {
				tt.on(f)
			}

			if err := a.CancelBasket(tt.args.ctx, tt.args.cancel); (err != nil) != tt.wantErr {
				t.Errorf("CancelBasket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApplication_CheckoutBasket(t *testing.T) {
	store := &domain.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := &domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	item := domain.Item{
		StoreID:      store.ID,
		ProductID:    product.ID,
		StoreName:    store.Name,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     10,
	}
	type fields struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}
	type args struct {
		ctx      context.Context
		checkout CheckoutBasket
	}
	tests := map[string]struct {
		args    args
		on      func(f fields)
		wantErr bool
	}{
		"CheckoutBasket": {
			args: args{
				ctx: context.Background(),
				checkout: CheckoutBasket{
					ID:        "basket-id",
					PaymentID: "payment-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items: map[string]domain.Item{
						product.ID: item,
					},
					Status: domain.BasketIsOpen,
				}, nil)
				f.baskets.On("Save", context.Background(), mock.AnythingOfType("*domain.Basket")).Return(nil)
				f.publisher.On("Publish", context.Background(), mock.AnythingOfType("ddd.event")).Return(nil)
			},
		},
		"CheckoutBasket.NoItems": {
			args: args{
				ctx: context.Background(),
				checkout: CheckoutBasket{
					ID:        "basket-id",
					PaymentID: "payment-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsOpen,
				}, nil)
			},
			wantErr: true,
		},
		"CheckoutBasket.CheckedOut": {
			args: args{
				ctx: context.Background(),
				checkout: CheckoutBasket{
					ID:        "basket-id",
					PaymentID: "payment-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsCheckedOut,
				}, nil)
			},
			wantErr: true,
		},
		"CheckoutBasket.Canceled": {
			args: args{
				ctx: context.Background(),
				checkout: CheckoutBasket{
					ID:        "basket-id",
					PaymentID: "payment-id",
				},
			},
			on: func(f fields) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&domain.Basket{
					Aggregate:  es.NewAggregate("basket-id", domain.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "",
					Items:      make(map[string]domain.Item),
					Status:     domain.BasketIsCanceled,
				}, nil)
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f := fields{
				baskets:   domain.NewMockBasketRepository(t),
				stores:    domain.NewMockStoreRepository(t),
				products:  domain.NewMockProductRepository(t),
				publisher: ddd.NewMockEventPublisher[ddd.Event](t),
			}
			a := Application{
				baskets:   f.baskets,
				stores:    f.stores,
				products:  f.products,
				publisher: f.publisher,
			}
			if tt.on != nil {
				tt.on(f)
			}

			if err := a.CheckoutBasket(tt.args.ctx, tt.args.checkout); (err != nil) != tt.wantErr {
				t.Errorf("CheckoutBasket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func TestApplication_GetBasket(t *testing.T) {
// 	type fields struct {
// 		baskets   domain.BasketRepository
// 		stores    domain.StoreRepository
// 		products  domain.ProductRepository
// 		publisher ddd.EventPublisher[ddd.Event]
// 	}
// 	type args struct {
// 		ctx context.Context
// 		get GetBasket
// 	}
// 	tests := map[string]struct {
// 		fields  fields
// 		args    args
// 		want    *domain.Basket
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			a := Application{
// 				baskets:   tt.fields.baskets,
// 				stores:    tt.fields.stores,
// 				products:  tt.fields.products,
// 				publisher: tt.fields.publisher,
// 			}
// 			got, err := a.GetBasket(tt.args.ctx, tt.args.get)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetBasket() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetBasket() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestApplication_RemoveItem(t *testing.T) {
// 	type fields struct {
// 		baskets   domain.BasketRepository
// 		stores    domain.StoreRepository
// 		products  domain.ProductRepository
// 		publisher ddd.EventPublisher[ddd.Event]
// 	}
// 	type args struct {
// 		ctx    context.Context
// 		remove RemoveItem
// 	}
// 	tests := map[string]struct {
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			a := Application{
// 				baskets:   tt.fields.baskets,
// 				stores:    tt.fields.stores,
// 				products:  tt.fields.products,
// 				publisher: tt.fields.publisher,
// 			}
// 			if err := a.RemoveItem(tt.args.ctx, tt.args.remove); (err != nil) != tt.wantErr {
// 				t.Errorf("RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
//
// func TestApplication_StartBasket(t *testing.T) {
// 	type fields struct {
// 		baskets   domain.BasketRepository
// 		stores    domain.StoreRepository
// 		products  domain.ProductRepository
// 		publisher ddd.EventPublisher[ddd.Event]
// 	}
// 	type args struct {
// 		ctx   context.Context
// 		start StartBasket
// 	}
// 	tests := map[string]struct {
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			a := Application{
// 				baskets:   tt.fields.baskets,
// 				stores:    tt.fields.stores,
// 				products:  tt.fields.products,
// 				publisher: tt.fields.publisher,
// 			}
// 			if err := a.StartBasket(tt.args.ctx, tt.args.start); (err != nil) != tt.wantErr {
// 				t.Errorf("StartBasket() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
//
// func TestNew(t *testing.T) {
// 	type args struct {
// 		baskets   domain.BasketRepository
// 		stores    domain.StoreRepository
// 		products  domain.ProductRepository
// 		publisher ddd.EventPublisher[ddd.Event]
// 	}
// 	tests := map[string]struct {
// 		name string
// 		args args
// 		want *Application
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			if got := New(tt.args.baskets, tt.args.stores, tt.args.products, tt.args.publisher); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("New() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
