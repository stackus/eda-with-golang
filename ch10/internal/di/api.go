package di

import (
	"context"
)

type contextKey int

const containerKey contextKey = 1

func Get(ctx context.Context, key string) any {
	ctn, ok := ctx.Value(containerKey).(*container)
	if !ok {
		panic("container does not exist on context")
	}

	return ctn.Get(key)
}
