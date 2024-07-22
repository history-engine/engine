package localcache

import (
	"history-engine/engine/library/shutdown"
)

var _ shutdown.Stopper = stopper{}

type stopper struct{}

func initSignal() {
	shutdown.RegisterStopper("local-cache", stopper{})
}

func (m stopper) Stop() error {
	persistent()
	return nil
}
