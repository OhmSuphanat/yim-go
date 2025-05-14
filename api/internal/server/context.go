package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ContextWithSignal(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stopChan
		cancel()
	}()

	return ctx
}
