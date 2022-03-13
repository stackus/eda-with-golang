package codecs

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type ProtoCodec struct {
	r registry.Registry
}

var _ registry.Codec = (*ProtoCodec)(nil)

func NewProtoCodec(r registry.Registry) *ProtoCodec {
	return &ProtoCodec{r: r}
}

func (c ProtoCodec) Register(name string, v interface{}, options ...registry.BuildOption) error {
	if _, ok := v.(proto.Message); !ok {
		return fmt.Errorf("%s does not implement proto.Message", name)
	}
	return registry.Register(c.r, name, v, c.marshal, c.unmarshal, options)
}

func (c ProtoCodec) RegisterFactory(name string, fn func() interface{}, options ...registry.BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("%s factory returns a nil value", name)
	} else if _, ok := v.(proto.Message); !ok {
		return fmt.Errorf("%s does not implement proto.Message", name)
	}
	return registry.RegisterFactory(c.r, name, fn, c.marshal, c.unmarshal, options)
}

func (ProtoCodec) marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (ProtoCodec) unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
