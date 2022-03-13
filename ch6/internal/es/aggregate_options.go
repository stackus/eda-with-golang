package es

import (
	"fmt"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type Versioner interface {
	setVersion(int)
}

func SetVersion(version int) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(Versioner); ok {
			agg.setVersion(version)
			return nil
		}
		return fmt.Errorf("%T does not implement setVersion(int)", v)
	}
}
