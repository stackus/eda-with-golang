package registry

type Codec interface {
	Register(name string, v interface{}, options ...BuildOption) error
	RegisterFactory(name string, fn func() interface{}, options ...BuildOption) error
}
