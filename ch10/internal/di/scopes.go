package di

type Scope int

const (
	Singleton Scope = iota + 1
	Scoped
	Transient
)
