package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/stackus/eda-with-golang/ch4/internal/config"
	"github.com/stackus/eda-with-golang/ch4/internal/egress"
	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type app struct {
	cfg     config.AppConfig
	modules []monolith.Module
	db      *sql.DB
	mux     *chi.Mux
	rpc     *grpc.Server
	waiter  egress.Waiter
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) RPC() *grpc.Server {
	return a.rpc
}

func (a *app) Wait(fns ...egress.WaitFunc) error {
	return a.waiter.Wait(fns...)
}

func (a *app) Context() context.Context {
	return a.waiter.Context()
}

func (a *app) CancelFunc() context.CancelFunc {
	return a.waiter.CancelFunc()
}

func (a *app) startupModules() error {
	for _, module := range a.modules {
		if err := module.Startup(a.Context(), a); err != nil {
			return err
		}
	}

	return nil
}

func (a *app) waitForWeb(ctx context.Context) error {
	webServer := &http.Server{
		Addr:    a.cfg.Web.Address(),
		Handler: a.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}

func (a *app) waitForRPC(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.cfg.Rpc.Address())
	if err != nil {
		return err
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := a.RPC().Serve(listener); err != nil && err != grpc.ErrServerStopped {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		stopped := make(chan struct{})
		go func() {
			a.RPC().GracefulStop()
			close(stopped)
		}()
		timeout := time.NewTimer(a.cfg.ShutdownTimeout)
		select {
		case <-timeout.C:
			// Force it to stop
			a.RPC().Stop()
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return group.Wait()
}
