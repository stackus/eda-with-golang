package codecs

import (
	"google.golang.org/protobuf/proto"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type ProtoCodec struct {
	r registry.Registry
}

func NewProtoCodec(r registry.Registry) *ProtoCodec {
	return &ProtoCodec{r: r}
}

func (c ProtoCodec) Register(name string, v interface{}, options ...registry.UnmarshalOption) error {
	return c.r.Register(name, v, c.marshal, c.unmarshal, options)
}

func (c ProtoCodec) RegisterFactory(name string, fn func() interface{}, options ...registry.UnmarshalOption) error {
	return c.r.RegisterFactory(name, fn, c.marshal, c.unmarshal, options)
}

func (ProtoCodec) marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (ProtoCodec) unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
