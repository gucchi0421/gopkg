package app

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdownSignal() <-chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}
