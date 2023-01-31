package useCases

import (
	"context"
	"github.com/ssoql/faq-service/internal/global"
	"os"
)

type testShutdownSig struct{}

func (testShutdownSig) String() string {
	return "shutdown signal"
}

func (testShutdownSig) Signal() {}

// handleShutdown returns context that will be cancelled if shutdown will occur
func handleShutdown(ctx context.Context) context.Context {
	shutdownChan := ctx.Value(global.ShutdownSignal).(chan os.Signal)
	shutdownCtx, cancel := context.WithCancel(ctx)

	go func(cancel context.CancelFunc) {
		<-shutdownChan
		cancel()
	}(cancel)

	return shutdownCtx
}
