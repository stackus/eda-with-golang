package em

type PublisherOption interface {
	configurePublisherConfig(*PublisherConfig)
}

type PublisherConfig struct {
	headers Headers
}

func NewPublisherConfig(options ...PublisherOption) PublisherConfig {
	cfg := PublisherConfig{
		headers: make(Headers),
	}

	for _, option := range options {
		option.configurePublisherConfig(&cfg)
	}

	return cfg
}

func (c PublisherConfig) Headers() Headers {
	return c.headers
}

func (h Headers) configurePublisherConfig(cfg *PublisherConfig) {
	cfg.headers = h
}
