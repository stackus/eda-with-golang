package registry

import (
	"fmt"
	"sync"
)

type (
	Registerable interface {
		Key() string
	}

	Marshaller   func(v interface{}) ([]byte, error)
	Unmarshaller func(d []byte, v interface{}) error

	Registry interface {
		Marshal(key string, v interface{}) ([]byte, error)
		Build(key string, options ...BuildOption) (interface{}, error)
		Unmarshal(key string, data []byte, options ...BuildOption) (interface{}, error)
		register(key string, fn func() interface{}, m Marshaller, u Unmarshaller, o []BuildOption) error
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

func (r *registry) Marshal(key string, v interface{}) ([]byte, error) {
	reg, exists := r.registered[key]
	if !exists {
		return nil, fmt.Errorf("nothing has been registered with the key `%s`", key)
	}
	return reg.marshaller(v)
}

func (r *registry) Unmarshal(key string, data []byte, options ...BuildOption) (interface{}, error) {
	v, err := r.Build(key, options...)
	if err != nil {
		return nil, err
	}

	err = r.registered[key].unmarshaller(data, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *registry) Build(key string, options ...BuildOption) (interface{}, error) {
	reg, exists := r.registered[key]
	if !exists {
		return nil, fmt.Errorf("nothing has been registered with the key `%s`", key)
	}

	v := reg.factory()
	uos := append(r.registered[key].options, options...)

	for _, option := range uos {
		err := option(v)
		if err != nil {
			return nil, err
		}
	}

	return v, nil
}

func (r *registry) register(key string, fn func() interface{}, m Marshaller, u Unmarshaller, o []BuildOption,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.registered[key]; exists {
		return fmt.Errorf("something with the key `%s` has already been registered", key)
	}

	r.registered[key] = registered{
		factory:      fn,
		marshaller:   m,
		unmarshaller: u,
		options:      o,
	}

	return nil
}
