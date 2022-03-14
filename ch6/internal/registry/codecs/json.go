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

func (c JSONCodec) Register(v registry.Registerable, options ...registry.BuildOption) error {
	return registry.Register(c.r, v, c.marshal, c.unmarshal, options)
}

func (c JSONCodec) RegisterKey(key string, v interface{}, options ...registry.BuildOption) error {
	return registry.RegisterKey(c.r, key, v, c.marshal, c.unmarshal, options)
}

func (c JSONCodec) RegisterFactory(key string, fn func() interface{}, options ...registry.BuildOption) error {
	return registry.RegisterFactory(c.r, key, fn, c.marshal, c.unmarshal, options)
}

func (JSONCodec) marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONCodec) unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
