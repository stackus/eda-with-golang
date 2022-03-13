package registry

import (
	"fmt"
	"sync"
)

type (
	Marshaller   func(v interface{}) ([]byte, error)
	Unmarshaller func(d []byte, v interface{}) error

	BuildOption func(v interface{}) error

	Registry interface {
		Build(name string, options ...BuildOption) (interface{}, error)
		Marshal(name string, v interface{}) ([]byte, error)
		Unmarshal(name string, data []byte, options ...BuildOption) (interface{}, error)
		register(name string, fn func() interface{}, m Marshaller, u Unmarshaller, o []BuildOption) error
	}
)

type registered struct {
	factory      func() interface{}
	marshaller   Marshaller
	unmarshaller Unmarshaller
	options      []BuildOption
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

func (r *registry) Marshal(name string, v interface{}) ([]byte, error) {
	if reg, exists := r.registered[name]; !exists {
		return nil, fmt.Errorf("nothing has been registered with the name `%s`", name)
	} else {
		return reg.marshaller(v)
	}
}

func (r *registry) Unmarshal(name string, data []byte, options ...BuildOption) (interface{}, error) {
	v, err := r.Build(name, options...)
	if err != nil {
		return nil, err
	}

	err = r.registered[name].unmarshaller(data, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *registry) Build(name string, options ...BuildOption) (interface{}, error) {
	if reg, exists := r.registered[name]; !exists {
		return nil, fmt.Errorf("nothing has been registered with the name `%s`", name)
	} else {
		v := reg.factory()
		uos := append(r.registered[name].options, options...)

		for _, option := range uos {
			err := option(v)
			if err != nil {
				return nil, err
			}
		}

		return v, nil
	}
}

func (r *registry) register(name string, fn func() interface{}, m Marshaller, u Unmarshaller, o []BuildOption,
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
