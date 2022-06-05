package di

import (
	"context"
	"fmt"
	"sync"
)

type DependencyFactoryFunc func(c Container) (any, error)
type DependencyCleanupFunc func(v any) error

type Container interface {
	AddSingleton(key string, fn DependencyFactoryFunc, options ...DependencyOption)
	AddScoped(key string, fn DependencyFactoryFunc, options ...DependencyOption)
	AddTransient(key string, fn DependencyFactoryFunc, options ...DependencyOption)
	Scoped(ctx context.Context) (context.Context, func() error)
	Get(key string) any
	Cleanup() error
}

type dependencyInfo struct {
	key     string
	scope   Scope
	factory DependencyFactoryFunc
	cleanup DependencyCleanupFunc
}

var _ Container = (*container)(nil)

type container struct {
	parent *container
	deps   map[string]dependencyInfo
	vals   map[string]any
	seen   seen
	mu     sync.Mutex
}

func New() Container {
	return &container{
		deps: make(map[string]dependencyInfo),
		vals: make(map[string]any),
	}
}

func (c *container) AddSingleton(key string, fn DependencyFactoryFunc, options ...DependencyOption) {
	i := dependencyInfo{
		key:     key,
		scope:   Singleton,
		factory: fn,
	}

	for _, option := range options {
		option.configureDependency(&i)
	}

	c.deps[key] = i
}

func (c *container) AddScoped(key string, fn DependencyFactoryFunc, options ...DependencyOption) {
	i := dependencyInfo{
		key:     key,
		scope:   Scoped,
		factory: fn,
	}

	for _, option := range options {
		option.configureDependency(&i)
	}

	c.deps[key] = i
}

func (c *container) AddTransient(key string, fn DependencyFactoryFunc, options ...DependencyOption) {
	i := dependencyInfo{
		key:     key,
		scope:   Transient,
		factory: fn,
	}

	for _, option := range options {
		option.configureDependency(&i)
	}

	c.deps[key] = i
}

func (c *container) Scoped(ctx context.Context) (context.Context, func() error) {
	newContainer := c.child()

	newCtx := context.WithValue(ctx, containerKey, newContainer)

	return newCtx, func() error {
		return newContainer.Cleanup()
	}
}

func (c *container) Get(key string) any {
	info, exists := c.deps[key]
	if !exists {
		panic(fmt.Sprintf("there is no dependency registered with `%s`", key))
	}

	// catch building Foo needs Bar and building Bar needs Foo :boom:
	if _, exists := c.seen[info.key]; exists {
		panic(fmt.Sprintf("cyclic dependencies encountered while building `%s`, built: %s", info.key, c.seen))
	}

	switch info.scope {
	case Transient:
		// get and return
		return c.build(info)
	case Singleton:
		// try to get it from the parent container
		return c.getFromParent(info)
	default:
		return c.get(info)
		// try to get it from this container
	}
}

func (c *container) Cleanup() error {
	// TODO track creation order... cleanup in reverse order
	for k, v := range c.vals {
		info, exists := c.deps[k]
		if !exists {
			// TODO don't do this
			continue
		}
		if info.cleanup != nil {
			err := info.cleanup(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *container) child() *container {
	return &container{
		parent: c,
		deps:   c.deps,
		vals:   make(map[string]any),
	}
}

func (c *container) getFromParent(info dependencyInfo) any {
	if c.parent != nil {
		return c.parent.getFromParent(info)
	}
	return c.get(info)
}

func (c *container) get(info dependencyInfo) any {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, exists := c.vals[info.key]; exists {
		return v
	}

	v := c.build(info)

	c.vals[info.key] = v

	return v
}

func (c *container) build(info dependencyInfo) any {
	v, err := info.factory(c.builder(info))
	if err != nil {
		panic(fmt.Sprintf("error building dependency `%s`: %s", info.key, err))
	}

	return v
}

func (c *container) builder(info dependencyInfo) *container {
	return &container{
		parent: c.parent,
		deps:   c.deps,
		vals:   c.vals,
		seen:   c.seen.add(info),
	}
}

// // NewContainer() *Container
// // BuildContext(ctx, cntr) ctx
// // di.
//
// func NewContainerContext(ctx context.Context) context.Context {
// 	ctnr := &container{}
// 	return context.WithValue(ctx, containerKey, ctnr)
// }
//
// func (c *container) get(key string) any {
// 	v, err := c.mayGet(key)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return v
// }
//
// func (c *container) mayGet(key string) (any, error) {
//
// }

/*


c := di.NewContainer()

c.AddSingleton("registry", func(c di.Container) (any, error) {
	reg := registry.New()

	// init the registry

	return reg, nil
})
c.AddSingleton("db", func(c di.Container) (any, error) {
	return sql.Open("pgx", cfg.PG.Conn)
})
c.AddScoped("tx", func(c di.Container) (any, error) {
	return c.Get("db").(*sql.DB).StartTx()
}, func(v any, err error) {
	// commit/rollback
})
c.AddScoped("orderRepo", func(c di.Container) (any, error) {
	return pg.NewOrderRepository("ordering.orders", c.Get("tx").(*sql.Tx), c.Get("registry").(registry.Registry))
})
c.AddScoped("app", func(c di.Container) (any, error) {
	return logging.Application(
		application.New(c.Get("orderRepo").(domain.OrderRepository),
		"Application", c.Get("logger"),
	)
})


c.StartScope(ctx, func(ctx, c) error {
	app := c.Get("app").(application.App)

	err := app.???(ctx, ...)


})



*/

/*

- get "thing"
- is "thing" already created?
	- if so then return "thing"
	- if not then call "thing" factory
		- if err panic(err)
		- if "thing" is not transient store "thing"


*/
