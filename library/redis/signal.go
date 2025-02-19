package redis

import (
	"history-engine/engine/library/shutdown"
)

var _ shutdown.Stopper = stopper{}

type stopper struct{}

func initSignal() {
	shutdown.RegisterStopper("redis", stopper{})
}

func (m stopper) Stop() error {
	// todo
	return nil
}
