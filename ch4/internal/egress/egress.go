package egress

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type WaitFunc func(ctx context.Context) error

type Waiter interface {
	Wait(fns ...WaitFunc) error
	Context() context.Context
	CancelFunc() context.CancelFunc
}

type waiter struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type egressCfg struct {
	parentCtx    context.Context
	catchSignals bool
}

func New(options ...EgressOption) Waiter {
	cfg := &egressCfg{
		parentCtx:    context.Background(),
		catchSignals: false,
	}

	for _, option := range options {
		option(cfg)
	}

	w := waiter{}
	w.ctx, w.cancel = context.WithCancel(cfg.parentCtx)
	if cfg.catchSignals {
		w.ctx, _ = signal.NotifyContext(w.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	return w
}

func (w waiter) Wait(fns ...WaitFunc) (err error) {
	errc := make(chan error)
	for _, fn := range fns {
		fn := fn
		go func() { errc <- fn(w.ctx) }()
	}
	for range fns {
		if err != nil {
			<-errc
		} else {
			err = <-errc
		}
		w.cancel()
	}
	return
}

func (w waiter) Context() context.Context {
	return w.ctx
}

func (w waiter) CancelFunc() context.CancelFunc {
	return w.cancel
}
