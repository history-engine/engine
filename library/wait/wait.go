package wait

import (
	"sync"
)

// 组件加载完才能启动服务器
var wg = sync.WaitGroup{}

func AddWait(delta int) {
	wg.Add(delta)
}

func Done() {
	wg.Done()
}

func Wait() {
	wg.Wait()
}
