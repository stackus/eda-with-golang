package ddd

import (
	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

func SetAggregateID(id string) registry.UnmarshalOption {
	return func(v interface{}) error {
		if agg, ok := v.(interface{ setID(string) }); ok {
			agg.setID(id)
		}
		return nil
	}
}

func SetAggregateName(name string) registry.UnmarshalOption {
	return func(v interface{}) error {
		if agg, ok := v.(interface{ setAggregateName(string) }); ok {
			agg.setAggregateName(name)
		}
		return nil
	}
}
