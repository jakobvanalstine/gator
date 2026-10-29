package isignal

import (
	"context"
	"os/signal"
)

// NotifyBackground is a wrapper around signal.NotifyContext.
// It is intended for creating top-level background contexts.
// The signals it passes to signal.NotifyContext are defined
// in files guarded by build tags.
func NotifyBackground() (background context.Context, stop context.CancelFunc) {
	return signal.NotifyContext(context.Background(), StopSignals...)
}
