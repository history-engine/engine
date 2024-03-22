package shutdown

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var stopMap = make(map[string]Stopper)

// Stopper 组件关闭回调
type Stopper interface {
	Stop() error
}

// RegisterStopper 注册关闭组件回调
func RegisterStopper(name string, stopper Stopper) {
	stopMap[name] = stopper
}

// ShutdownComponent 关闭组件
func ShutdownComponent(ch chan error) {
	log.Println("shutdown component")

	timeout := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(timeout)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(len(stopMap))
	for name, stopper := range stopMap {
		go func(wg *sync.WaitGroup, name string, stopper Stopper) {
			if err := stopper.Stop(); err != nil {
				ch <- fmt.Errorf("stop %s err: %v", name, err)
			}
			wg.Done()
		}(wg, name, stopper)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:

	case <-timeout:
		wg.Done()
	}
}
