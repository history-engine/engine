package page

import (
	"history-engine/engine/ent"
	"history-engine/engine/model"
	"history-engine/engine/utils"
	"strings"
)

func Filter(row ent.Page) (bool, error) {
	// todo 把过滤方式：文件后缀、host、大小、频率等封装起来

	return true, nil
}

func ParseHtmlInfo(filename string) *model.HtmlInfo {
	if len(filename) == 0 {
		return nil
	}

	sff := &model.HtmlInfo{}
	kvs := strings.Split(filename, "_")
	for _, item := range kvs {
		kv := strings.SplitN(item, "-", 2)
		if len(kv) != 2 {
			continue
		}

		switch kv[0] {
		case "host":
			sff.Host = kv[1]
		case "suffix":
			sff.Suffix = utils.FileSuffix(kv[1])
		case "sha1":
			sff.Sha1 = kv[1]
		}
	}

	if sff.Sha1 == "" && sff.Url != "" {
		sff.Sha1 = utils.Sha1Str(sff.Sha1)
	}

	if sff.Sha1 == "" {
		return nil
	}

	return sff
}
