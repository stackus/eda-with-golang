package codecs

import (
	"encoding/json"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type JSONCodec struct {
	r registry.Registry
}

var _ registry.Codec = (*JSONCodec)(nil)

func NewJSONCodec(r registry.Registry) *JSONCodec {
	return &JSONCodec{r: r}
}

func (c JSONCodec) Register(name string, v interface{}, options ...registry.BuildOption) error {
	return registry.Register(c.r, name, v, c.marshal, c.unmarshal, options)
}

func (c JSONCodec) RegisterFactory(name string, fn func() interface{}, options ...registry.BuildOption) error {
	return registry.RegisterFactory(c.r, name, fn, c.marshal, c.unmarshal, options)
}

func (JSONCodec) marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONCodec) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
