package codecs

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/stackus/eda-with-golang/ch6/internal/registry"
)

type ProtoCodec struct {
	r registry.Registry
}

var _ registry.Codec = (*ProtoCodec)(nil)
var protoT = reflect.TypeOf((*proto.Message)(nil)).Elem()

func NewProtoCodec(r registry.Registry) *ProtoCodec {
	return &ProtoCodec{r: r}
}

func (c ProtoCodec) Register(v registry.Registerable, options ...registry.BuildOption) error {
	if !reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.Register(c.r, v, c.marshal, c.unmarshal, options)
}

func (c ProtoCodec) RegisterKey(key string, v interface{}, options ...registry.BuildOption) error {
	if !reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.RegisterKey(c.r, key, v, c.marshal, c.unmarshal, options)
}

func (c ProtoCodec) RegisterFactory(key string, fn func() interface{}, options ...registry.BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("%s factory returns a nil value", key)
	} else if _, ok := v.(proto.Message); !ok {
		return fmt.Errorf("%s does not implement proto.Message", key)
	}
	return registry.RegisterFactory(c.r, key, fn, c.marshal, c.unmarshal, options)
}

func (ProtoCodec) marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (ProtoCodec) unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
