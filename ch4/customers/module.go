package customers

import (
	"context"

	"github.com/stackus/eda-with-golang/ch4/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	return nil
}
