package search

import (
	"history-engine/engine/setting"
	"sync"
)

var engine EngineInterface
var once sync.Once

func Engine() EngineInterface {
	once.Do(func() {
		switch setting.Search.Engine {
		case "zinc":
			engine = NewZincSearch()
		case "meili":
			engine = NewMeiliSearch()
		default:
			panic("wrong search engine")
		}
	})
	return engine
}
