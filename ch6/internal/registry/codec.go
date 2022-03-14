package registry

type Codec interface {
	Register(v Registerable, options ...BuildOption) error
	RegisterKey(key string, v interface{}, options ...BuildOption) error
	RegisterFactory(key string, fn func() interface{}, options ...BuildOption) error
}
