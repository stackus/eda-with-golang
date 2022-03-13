package ddd

import (
	"fmt"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type IDer interface {
	setID(string)
}

func SetID(id string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(IDer); ok {
			e.setID(id)
			return nil
		}
		return fmt.Errorf("%T does not implement setID(string)", v)
	}
}

type Namer interface {
	setName(string)
}

func SetName(name string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(Namer); ok {
			e.setName(name)
			return nil
		}
		return fmt.Errorf("%T does not implement setName(string)", v)
	}
}
