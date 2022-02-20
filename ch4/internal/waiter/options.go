package waiter

import (
	"context"
)

type EgressOption func(c *egressCfg)

func ParentContext(ctx context.Context) EgressOption {
	return func(c *egressCfg) {
		c.parentCtx = ctx
	}
}

func CatchSignals() EgressOption {
	return func(c *egressCfg) {
		c.catchSignals = true
	}
}
