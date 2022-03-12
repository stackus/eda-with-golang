package registry

import (
	"fmt"
	"reflect"
	"sync"
)

type (
	Marshaller   func(v interface{}) ([]byte, error)
	Unmarshaller func(d []byte, v interface{}) error

	UnmarshalOption func(v interface{}) error

	Registry interface {
		Register(name string, v interface{}, m Marshaller, u Unmarshaller, options []UnmarshalOption) error
		RegisterFactory(name string, fn func() interface{}, m Marshaller, u Unmarshaller, options []UnmarshalOption,
		) error
		Marshal(name string, v interface{}) ([]byte, error)
		Unmarshal(name string, data []byte, options ...UnmarshalOption) (interface{}, error)
	}
)

type registered struct {
	factory      func() interface{}
	marshaller   Marshaller
	unmarshaller Unmarshaller
	options      []UnmarshalOption
}

type registry struct {
	registered map[string]registered
	mu         sync.RWMutex
}

func New() *registry {
	return &registry{
		registered: make(map[string]registered),
	}
}

func (r *registry) Register(name string, v interface{}, m Marshaller, u Unmarshaller, os []UnmarshalOption) error {
	t := reflect.TypeOf(v)

	// (*T)(nil), *T{}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return r.register(name, func() interface{} {
		return reflect.New(t).Interface()
	}, m, u, os)
}

func (r *registry) RegisterFactory(name string, fn func() interface{}, m Marshaller, u Unmarshaller,
	os []UnmarshalOption,
) error {
	if v := fn(); v == nil {
		return fmt.Errorf("factory for item `%s` returns a nil value", name)
	}

	if t := reflect.TypeOf(fn()); t.Kind() != reflect.Ptr {
		return fmt.Errorf("factory for item `%s` does not return a pointer receiver", name)
	}

	return r.register(name, fn, m, u, os)
}

func (r *registry) Marshal(name string, v interface{}) ([]byte, error) {
	if reg, exists := r.registered[name]; !exists {
		return nil, fmt.Errorf("nothing has been registered with the name `%s`", name)
	} else {
		return reg.marshaller(v)
	}
}

func (r *registry) Unmarshal(name string, data []byte, options ...UnmarshalOption) (interface{}, error) {
	if reg, exists := r.registered[name]; !exists {
		return nil, fmt.Errorf("nothing has been registered with the name `%s`", name)
	} else {
		v := reg.factory()
		fmt.Printf("TYPE: %v %T\n\n", v, v)
		err := reg.unmarshaller(data, v)
		if err != nil {
			return nil, err
		}

		uos := append(r.registered[name].options, options...)

		for _, option := range uos {
			err = option(v)
			if err != nil {
				return nil, err
			}
		}

		return v, err
	}
}

func (r *registry) register(name string, fn func() interface{}, m Marshaller, u Unmarshaller, o []UnmarshalOption,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.registered[name]; exists {
		return fmt.Errorf("something with the name `%s` has already been registered", name)
	}

	r.registered[name] = registered{
		factory:      fn,
		marshaller:   m,
		unmarshaller: u,
		options:      o,
	}
	return nil
}

func WithSetID(id string) UnmarshalOption {
	return WithSetStringField("ID", id)
}

func WithSetVersion(version int) UnmarshalOption {
	return func(v interface{}) error {
		if agg, ok := v.(interface{ SetVersion(version int) }); ok {
			agg.SetVersion(version)
		}
		return nil
	}
}

func WithSetStringField(field, value string) UnmarshalOption {
	return func(v interface{}) error {
		p := reflect.ValueOf(v)
		e := p.Elem()
		if e.Kind() == reflect.Struct {
			f := e.FieldByName(field)
			if f.IsValid() && f.CanSet() {
				if f.Kind() == reflect.String {
					f.SetString(value)
				}
			}
		}
		return nil
	}
}

func WithSetIntField(field string, value int) UnmarshalOption {
	return func(v interface{}) error {
		p := reflect.ValueOf(v)
		e := p.Elem()
		if e.Kind() == reflect.Struct {
			f := e.FieldByName(field)
			if f.IsValid() && f.CanSet() {
				if f.Kind() == reflect.Int {
					x := int64(value)
					if !f.OverflowInt(x) {
						f.SetInt(x)
					}
				}
			}
		}
		return nil
	}
}
