package registry

import (
	"fmt"
	"reflect"
)

func Register(reg Registry, name string, v interface{}, m Marshaller, u Unmarshaller, os []BuildOption) error {
	t := reflect.TypeOf(v)

	// accept (*T)(nil), *T{}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reg.register(name, func() interface{} {
		return reflect.New(t).Interface()
	}, m, u, os)
}

func RegisterFactory(reg Registry, name string, fn func() interface{}, m Marshaller, u Unmarshaller,
	os []BuildOption,
) error {
	if v := fn(); v == nil {
		return fmt.Errorf("factory for item `%s` returns a nil value", name)
	}

	if t := reflect.TypeOf(fn()); t.Kind() != reflect.Ptr {
		return fmt.Errorf("factory for item `%s` does not return a pointer receiver", name)
	}

	return reg.register(name, fn, m, u, os)
}
