package page

import (
	"history-engine/engine/ent"
)

func Filter(row ent.Page) (bool, error) {
	// todo 把过滤方式：文件后缀、host、大小、频率等封装起来

	return true, nil
}
