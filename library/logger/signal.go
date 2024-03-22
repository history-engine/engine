package logger

import (
	"history-engine/engine/library/shutdown"
)

var _ shutdown.Stopper = stopper{}

type stopper struct{}

func initSignal() {
	shutdown.RegisterStopper("db", stopper{})
}

func (m stopper) Stop() error {
	// todo
	return nil
}
